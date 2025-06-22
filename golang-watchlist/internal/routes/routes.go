package routes

import (
	"database/sql"
	"encoding/json"
	"golang-watchlist/internal/models"
	"golang-watchlist/internal/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleRecordRoutes(router *mux.Router, db *sql.DB) {
	repo := repository.NewRecordRepository(db)
	router.HandleFunc("/watchlist", CreateRecord(repo)).Methods(http.MethodPost)
	router.HandleFunc("/watchlist", GetRecords(repo)).Methods(http.MethodGet)
	router.HandleFunc("/watchlist/{id}", UpdateRecord(repo)).Methods(http.MethodPut)
	router.HandleFunc("/watchlist/{id}", DeleteRecord(repo)).Methods(http.MethodDelete)
}

func CreateRecord(repo repository.RecordRepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record models.Record
		if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if record.Title == "" || record.Status == "" {
			http.Error(w, "Title and Status are required", http.StatusBadRequest)
			return
		}

		if err := repo.CreateRecord(r.Context(), &record); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(record)
	}
}

func GetRecords(repo repository.RecordRepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		records, err := repo.GetRecords(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(records)
	}
}

func UpdateRecord(repo repository.RecordRepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var record models.Record
		if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if record.Title == "" || record.Status == "" {
			http.Error(w, "Title and Status are required", http.StatusBadRequest)
			return
		}

		record.ID = id
		if err := repo.UpdateRecord(r.Context(), &record); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(record)
	}
}

func DeleteRecord(repo repository.RecordRepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := repo.DeleteRecord(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
