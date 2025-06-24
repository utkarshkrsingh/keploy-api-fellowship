package integration

import (
	"context"
	"database/sql"
	"golang-watchlist/internal/db"
	"golang-watchlist/internal/models"
	"golang-watchlist/internal/repository"
	"path/filepath"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func setupTestDB(t *testing.T) *sql.DB {
	envPath := filepath.Join("..", "..", ".env")
	if err := godotenv.Load(envPath); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}
	database, err := db.NewDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return database
}

func TestCreateAndGetRecord(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewRecordRepository(db)
	record := models.Record{
		Title:           "Bleach",
		TotalEpisodes:   366,
		WatchedEpisodes: 356,
		Type:            "TV",
		Status:          "watching",
	}

	err := repo.CreateRecord(context.Background(), &record)
	if err != nil {
		t.Fatalf("Failed to create record: %v", err)
	}
	if record.ID == 0 {
		t.Errorf("Expected ID to be set, got 0")
	}

	records, err := repo.GetRecords(context.Background())
	if err != nil {
		t.Fatalf("Failed to get records: %v", err)
	}
	if len(records) == 0 {
		t.Errorf("Expected at one book, get none")
	}
}

func TestUpdateRecord(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewRecordRepository(db)
	record := models.Record{
		Title:           "Naruto",
		TotalEpisodes:   220,
		WatchedEpisodes: 219,
		Type:            "TV",
		Status:          "watching",
	}

	err := repo.CreateRecord(context.Background(), &record)
	if err != nil {
		t.Fatalf("Failed to create record: %v", err)
	}

	record.WatchedEpisodes = 220
	record.Status = "completed"
	err = repo.UpdateRecord(context.Background(), &record)
	if err != nil {
		t.Fatalf("Failed to update record: %v", err)
	}

	records, err := repo.GetRecords(context.Background())
	if err != nil {
		t.Fatalf("Failed to get records: %v", err)
	}
	found := false
	for _, b := range records {
		if b.ID == record.ID && b.Title == record.Title && b.WatchedEpisodes == record.WatchedEpisodes {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected record not found")
	}
}

func TestDeleteRecord(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewRecordRepository(db)
	record := models.Record{
		Title:           "Boruto: Naruto Next Generations",
		TotalEpisodes:   293,
		WatchedEpisodes: 0,
		Type:            "TV",
		Status:          "planning",
	}

	err := repo.CreateRecord(context.Background(), &record)
	if err != nil {
		t.Fatalf("Failed to create record: %v", err)
	}

	err = repo.DeleteRecord(context.Background(), record.ID)
	if err != nil {
		t.Fatalf("Failed to delete record: %v", err)
	}

	records, err := repo.GetRecords(context.Background())
	if err != nil {
		t.Fatalf("Failed to get records: %v", err)
	}
	for _, b := range records {
		if b.ID == record.ID {
			t.Error("Expected record to be deleted")
		}
	}
}
