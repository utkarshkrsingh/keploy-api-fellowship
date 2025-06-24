package unit

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"golang-watchlist/internal/models"
	"golang-watchlist/internal/repository"
	"golang-watchlist/internal/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var testDB *sql.DB

// ====================================================================================================
// MOCK REPOSITORY FOR HANDLER TESTS
// ====================================================================================================

type mockRecordRepository struct {
	createFunc func(ctx context.Context, record *models.Record) error
	getFunc    func(ctx context.Context) ([]models.Record, error)
	updateFunc func(ctx context.Context, record *models.Record) error
	deleteFunc func(ctx context.Context, id int) error
}

func (m *mockRecordRepository) CreateRecord(ctx context.Context, record *models.Record) error {
	return m.createFunc(ctx, record)
}

func (m *mockRecordRepository) GetRecords(ctx context.Context) ([]models.Record, error) {
	return m.getFunc(ctx)
}

func (m *mockRecordRepository) UpdateRecord(ctx context.Context, record *models.Record) error {
	return m.updateFunc(ctx, record)
}

func (m *mockRecordRepository) DeleteRecord(ctx context.Context, id int) error {
	return m.deleteFunc(ctx, id)
}

var _ repository.RecordRepositoryInterface = &mockRecordRepository{}

// ====================================================================================================
// DATABASE SETUP FOR NON-MOCK TESTS
// ====================================================================================================

func TestMain(m *testing.M) {
	var err error
	testDB, err := setupTestDB()
	if err != nil {
		fmt.Printf("Failed to setup test database: %v\n", err)
		fmt.Println("Warning: Database tests will be skipped")
	} else if err := testDB.Ping(); err != nil {
		fmt.Printf("Cannot ping DB: %v\n", err)
		testDB.Close()
		testDB = nil
	}

	code := m.Run()

	if testDB != nil {
		tearDownTestDB(testDB)
	}
	os.Exit(code)
}

func setupTestDB() (*sql.DB, error) {
	dbUser := "root"
	dbPassword := "testqwerty"
	dbHost := "127.0.0.1"
	dbPort := "5011"
	dbName := "test-anime-watch-list"

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=UTF8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	testDB, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v", err)
	}

	if err := testDB.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping test database: %v", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS watch_list (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		total_episodes INTEGER,
		watched_episodes INTEGER,
		type TEXT,
		status TEXT NOT NULL
	)`
	_, err = testDB.Exec(schema)
	if err != nil {
		testDB.Close()
		return nil, fmt.Errorf("Failed to create watch_list table: %v", err)
	}

	return testDB, nil
}

func tearDownTestDB(db *sql.DB) {
	if db != nil {
		db.Exec("DROP TABLE IF EXISTS watch_list")
		db.Close()
	}
}

func skipIfNoDatabase(t *testing.T) {
	if testDB == nil {
		t.Skip("Skipping database test - database not variable")
	}
}

// ====================================================================================================
// MOCK TESTS (Handlers Tests)
// ====================================================================================================

func TestCreateRecord_Mock(t *testing.T) {
	tests := []struct {
		name           string
		inputRecord    models.Record
		createFunc     func(ctx context.Context, r *models.Record) error
		expectedStatus int
		expectedId     int
		expectedError  bool
	}{
		{
			name: "Successful Creation",
			inputRecord: models.Record{
				Title:           "Attack on Titan",
				TotalEpisodes:   25,
				WatchedEpisodes: 24,
				Type:            "TV",
				Status:          "watching",
			},
			createFunc: func(ctx context.Context, r *models.Record) error {
				r.ID = 1
				return nil
			},
			expectedStatus: http.StatusCreated,
			expectedId:     1,
			expectedError:  false,
		},
		{
			name: "Invalid request body - missing 'Title'",
			inputRecord: models.Record{
				Title:           "",
				TotalEpisodes:   12,
				WatchedEpisodes: 12,
				Type:            "TV",
				Status:          "completed",
			},
			createFunc:     func(ctx context.Context, r *models.Record) error { return nil },
			expectedStatus: http.StatusBadRequest,
			expectedId:     0,
			expectedError:  true,
		},
		{
			name:           "Invalid request body - missing body",
			inputRecord:    models.Record{},
			createFunc:     func(ctx context.Context, r *models.Record) error { return nil },
			expectedStatus: http.StatusBadRequest,
			expectedId:     0,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.inputRecord)
			req := httptest.NewRequest(http.MethodPost, "/watchlist", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			mockRepo := &mockRecordRepository{
				createFunc: tt.createFunc,
			}
			router := mux.NewRouter()
			router.HandleFunc("/watchlist", routes.CreateRecord(mockRepo)).Methods(http.MethodPost)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if !tt.expectedError {
				var createdRecord models.Record
				if err := json.NewDecoder(w.Body).Decode(&createdRecord); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if createdRecord.ID != tt.expectedId {
					t.Errorf("Expected record id %d, got %d", tt.expectedId, createdRecord.ID)
				}
				if createdRecord.Title != tt.inputRecord.Title {
					t.Errorf("Expected record title %s, got %s", tt.inputRecord.Title, createdRecord.Title)
				}
			}
		})
	}
}

func TestGetRecord_Mock(t *testing.T) {
	tests := []struct {
		name           string
		getFunc        func(ctx context.Context) ([]models.Record, error)
		expectedStatus int
		expectedRecord []models.Record
	}{
		{
			name: "Successful get",
			getFunc: func(ctx context.Context) ([]models.Record, error) {
				return []models.Record{
					{
						ID:              1,
						Title:           "Bleach",
						TotalEpisodes:   366,
						WatchedEpisodes: 366,
						Type:            "TV",
						Status:          "completed",
					},
				}, nil
			},
			expectedStatus: 200,
			expectedRecord: []models.Record{
				{
					ID:              1,
					Title:           "Bleach",
					TotalEpisodes:   366,
					WatchedEpisodes: 366,
					Type:            "TV",
					Status:          "completed",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/watchlist", nil)
			w := httptest.NewRecorder()

			mockRepo := &mockRecordRepository{
				getFunc: tt.getFunc,
			}
			router := mux.NewRouter()
			router.HandleFunc("/watchlist", routes.GetRecords(mockRepo)).Methods(http.MethodGet)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedRecord != nil {
				var records []models.Record
				if err := json.NewDecoder(w.Body).Decode(&records); err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}
				if len(records) != len(tt.expectedRecord) {
					t.Errorf("Expected %d records, got %d", len(tt.expectedRecord), len(records))
				}
			}
		})
	}
}

func TestUpdateRecord_Mock(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		inputRecord    models.Record
		updateFunc     func(ctx context.Context, record *models.Record) error
		expectedStatus int
		expectedErr    bool
	}{
		{
			name: "Successful Update",
			id:   "1",
			inputRecord: models.Record{
				Title:           "Bleach",
				TotalEpisodes:   366,
				WatchedEpisodes: 365,
				Type:            "TV",
				Status:          "watching",
			},
			updateFunc:     func(ctx context.Context, record *models.Record) error { return nil },
			expectedStatus: http.StatusOK,
			expectedErr:    false,
		},
		{
			name: "Invalid ID",
			id:   "invalid",
			inputRecord: models.Record{
				Title:           "Bleach",
				TotalEpisodes:   366,
				WatchedEpisodes: 366,
				Type:            "TV",
				Status:          "completed",
			},
			updateFunc:     func(ctx context.Context, record *models.Record) error { return nil },
			expectedStatus: http.StatusBadRequest,
			expectedErr:    true,
		},
		{
			name: "Repository Error",
			id:   "1",
			inputRecord: models.Record{
				Title:         "Bleach",
				TotalEpisodes: 366,
			},
			updateFunc: func(ctx context.Context, record *models.Record) error {
				return errors.New("database error")
			},
			expectedStatus: http.StatusBadRequest,
			expectedErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.inputRecord)
			req := httptest.NewRequest(http.MethodPut, "/watchlist/"+tt.id, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			mockRepo := &mockRecordRepository{
				updateFunc: tt.updateFunc,
			}

			router := mux.NewRouter()
			router.HandleFunc("/watchlist/{id}", routes.UpdateRecord(mockRepo)).Methods(http.MethodPut)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if !tt.expectedErr {
				var updatedRecord models.Record
				if err := json.NewDecoder(w.Body).Decode(&updatedRecord); err != nil {
					t.Fatalf("Failed to decode responce: %v", err)
				}

				if updatedRecord.Title != tt.inputRecord.Title {
					t.Errorf("Expected record title %s, got %s", tt.inputRecord.Title, updatedRecord.Title)
				}
			}
		})
	}
}

func TestDeleteRecord_Mock(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		deleteFunc     func(ctx context.Context, id int) error
		expectedStatus int
	}{
		{
			name:           "Successful Delete",
			id:             "1",
			deleteFunc:     func(ctx context.Context, id int) error { return nil },
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "Invalid ID",
			id:             "invalid",
			deleteFunc:     func(ctx context.Context, id int) error { return nil },
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Repository error",
			id:   "1",
			deleteFunc: func(ctx context.Context, id int) error {
				return errors.New("database error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/watchlist/"+tt.id, nil)
			w := httptest.NewRecorder()

			mockRepo := &mockRecordRepository{
				deleteFunc: tt.deleteFunc,
			}

			router := mux.NewRouter()
			router.HandleFunc("/watchlist/{id}", routes.DeleteRecord(mockRepo)).Methods(http.MethodDelete)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

// ====================================================================================================
// NON-MOCK TESTS (Repository Tests)
// ====================================================================================================

func TestCreateRecordRepository_Database(t *testing.T) {
	skipIfNoDatabase(t)

	repo := repository.NewRecordRepository(testDB)

	tests := []struct {
		name           string
		inputRecord    models.Record
		expectedStatus int
		expectedId     int
		expectedError  bool
	}{
		{
			name: "Successful Creation",
			inputRecord: models.Record{
				Title:           "Bleach",
				TotalEpisodes:   336,
				WatchedEpisodes: 0,
				Type:            "TV",
				Status:          "planning",
			},
			expectedStatus: http.StatusCreated,
			expectedId:     1,
			expectedError:  false,
		},
		{
			name: "Invalid request body - missing 'Title'",
			inputRecord: models.Record{
				Title:           "",
				TotalEpisodes:   12,
				WatchedEpisodes: 12,
				Type:            "TV",
				Status:          "completed",
			},
			expectedStatus: http.StatusBadRequest,
			expectedId:     0,
			expectedError:  true,
		},
		{
			name:           "Invalid request body - missing body",
			inputRecord:    models.Record{},
			expectedStatus: http.StatusBadRequest,
			expectedId:     0,
			expectedError:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.inputRecord)
			req := httptest.NewRequest(http.MethodPost, "/watchlist", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/watchlist", routes.CreateRecord(repo)).Methods(http.MethodPost)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if !tt.expectedError {
				var createdRecord models.Record
				if err := json.NewDecoder(w.Body).Decode(&createdRecord); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if createdRecord.ID != tt.expectedId {
					t.Errorf("Expected record id %d, got %d", tt.expectedId, createdRecord.ID)
				}
				if createdRecord.Title != tt.inputRecord.Title {
					t.Errorf("Expected record title %s, got %s", tt.inputRecord.Title, createdRecord.Title)
				}
			}
		})
	}
}

func TestGetRecord_Database(t *testing.T) {
	skipIfNoDatabase(t)

	repo := repository.NewRecordRepository(testDB)

	tests := []struct {
		name           string
		expectedStatus int
		expectedRecord []models.Record
	}{
		{
			name:           "Successful get",
			expectedStatus: http.StatusOK,
			expectedRecord: []models.Record{{ID: 1, Title: "Bleach", TotalEpisodes: 366, WatchedEpisodes: 366, Type: "TV", Status: "completed"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/watchlist", nil)
			w := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/watchlist", routes.GetRecords(repo)).Methods(http.MethodGet)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedRecord != nil {
				var records []models.Record
				if err := json.NewDecoder(w.Body).Decode(&records); err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}
				if len(records) != len(tt.expectedRecord) {
					t.Errorf("Expected %d records, got %d", len(tt.expectedRecord), len(records))
				}
			}
		})
	}
}

func TestUpdateRecord_Database(t *testing.T) {
	skipIfNoDatabase(t)

	repo := repository.NewRecordRepository(testDB)

	tests := []struct {
		name           string
		id             string
		inputRecord    models.Record
		expectedStatus int
		expectedErr    bool
	}{
		{
			name: "Successful Update",
			id:   "1",
			inputRecord: models.Record{
				Title:           "Bleach",
				TotalEpisodes:   366,
				WatchedEpisodes: 365,
				Type:            "TV",
				Status:          "watching",
			},
			expectedStatus: http.StatusOK,
			expectedErr:    false,
		},
		{
			name: "Invalid ID",
			id:   "invalid",
			inputRecord: models.Record{
				Title:           "Bleach",
				TotalEpisodes:   366,
				WatchedEpisodes: 366,
				Type:            "TV",
				Status:          "completed",
			},
			expectedStatus: http.StatusBadRequest,
			expectedErr:    true,
		},
		{
			name: "Repository Error",
			id:   "1",
			inputRecord: models.Record{
				Title:         "Bleach",
				TotalEpisodes: 366,
			},
			expectedStatus: http.StatusBadRequest,
			expectedErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.inputRecord)
			req := httptest.NewRequest(http.MethodPut, "/watchlist/"+tt.id, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/watchlist/{id}", routes.UpdateRecord(repo)).Methods(http.MethodPut)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, get %d", tt.expectedStatus, w.Code)
			}

			if !tt.expectedErr {
				var updatedRecord models.Record
				if err := json.NewDecoder(w.Body).Decode(&updatedRecord); err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}
				if tt.inputRecord.Title != updatedRecord.Title {
					t.Errorf("Expected record title %s, got %s", tt.inputRecord.Title, updatedRecord.Title)
				}
			}
		})
	}
}

func TestDeleteRecord_Database(t *testing.T) {
	skipIfNoDatabase(t)

	repo := repository.NewRecordRepository(testDB)

	tests := []struct {
		name           string
		id             string
		expectedStatus int
	}{
		{
			name:           "Successful Delete",
			id:             "1",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "Invalid ID",
			id:             "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Repository error",
			id:             "1",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/watchlist/"+tt.id, nil)
			w := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/watchlist/{id}", routes.DeleteRecord(repo)).Methods(http.MethodDelete)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
