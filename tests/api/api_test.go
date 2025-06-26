package api

import (
	"bytes"
	"encoding/json"
	"golang-watchlist/internal/db"
	"golang-watchlist/internal/models"
	"golang-watchlist/internal/routes"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func setupTestServer(t *testing.T) *mux.Router {
	envPath := filepath.Join("..", "..", ".env")
	if err := godotenv.Load(envPath); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}
	db, err := db.NewDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	router := mux.NewRouter()
	routes.HandleRecordRoutes(router, db)
	return router
}

func TestAPICreateAndGet(t *testing.T) {
	router := setupTestServer(t)

	record := models.Record{
		Title:           "JUJUTSU KAISEN",
		TotalEpisodes:   24,
		WatchedEpisodes: 24,
		Type:            "TV",
		Status:          "completed",
	}
	body, _ := json.Marshal(record)
	req := httptest.NewRequest(http.MethodPost, "/watchlist", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var createdRecord models.Record
	if err := json.NewDecoder(w.Body).Decode(&createdRecord); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	req = httptest.NewRequest(http.MethodGet, "/watchlist", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var records []models.Record
	if err := json.NewDecoder(w.Body).Decode(&records); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	if len(records) == 0 {
		t.Errorf("Expected at least one book, get none")
	}
}

func TestAPIUpdateBook(t *testing.T) {
	router := setupTestServer(t)

	record := models.Record{
		Title:           "Attack on Titan",
		TotalEpisodes:   25,
		WatchedEpisodes: 24,
		Type:            "TV",
		Status:          "watching",
	}
	body, _ := json.Marshal(record)
	req := httptest.NewRequest(http.MethodPost, "/watchlist", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var createdRecord models.Record
	if err := json.NewDecoder(w.Body).Decode(&createdRecord); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	updatedRecord := models.Record{
		Title:           "Attack on Titan",
		TotalEpisodes:   25,
		WatchedEpisodes: 25,
		Type:            "TV",
		Status:          "completed",
	}
	body, _ = json.Marshal(updatedRecord)
	req = httptest.NewRequest(http.MethodPut, "/watchlist/"+strconv.Itoa(createdRecord.ID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var resultRecord models.Record
	if err := json.NewDecoder(w.Body).Decode(&resultRecord); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	if resultRecord.Title != createdRecord.Title {
		t.Errorf("Expected record title %s, got %s", createdRecord.Title, resultRecord.Title)
	}
}

func TestAPIDeleteRecord(t *testing.T) {
	router := setupTestServer(t)

	record := models.Record{
		Title:           "Hunter x Hunter",
		TotalEpisodes:   148,
		WatchedEpisodes: 0,
		Type:            "TV",
		Status:          "planning",
	}

	body, _ := json.Marshal(record)
	req := httptest.NewRequest(http.MethodPost, "/watchlist", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var createdRecord models.Record
	if err := json.NewDecoder(w.Body).Decode(&createdRecord); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	req = httptest.NewRequest(http.MethodDelete, "/watchlist/"+strconv.Itoa(createdRecord.ID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status %d, got %d", http.StatusNoContent, w.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/watchlist", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var records []models.Record
	if err := json.NewDecoder(w.Body).Decode(&records); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	for _, b := range records {
		if b.ID == createdRecord.ID {
			t.Errorf("Expected record to be deleted")
		}
	}
}
