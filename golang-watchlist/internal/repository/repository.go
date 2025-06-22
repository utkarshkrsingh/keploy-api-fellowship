package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang-watchlist/internal/models"
)

type RecordRepositoryInterface interface {
	CreateRecord(ctx context.Context, record *models.Record) error
	GetRecords(ctx context.Context) ([]models.Record, error)
	UpdateRecord(ctx context.Context, record *models.Record) error
	DeleteRecord(ctx context.Context, id int) error
}

type RecordRepository struct {
	db *sql.DB
}

func NewRecordRepository(db *sql.DB) *RecordRepository {
	return &RecordRepository{db: db}
}

var _ RecordRepositoryInterface = &RecordRepository{}

func (r *RecordRepository) CreateRecord(ctx context.Context, record *models.Record) error {
	query := `
	INSERT INTO watch_list (title, total_episodes, watched_episodes, type, status)
	VALUES (?, ?, ?, ?, ?)
	`
	result, err := r.db.ExecContext(ctx, query,
		record.Title, record.TotalEpisodes,
		record.WatchedEpisodes, record.Type,
		record.Status)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	record.ID = int(id)
	return nil
}

func (r *RecordRepository) GetRecords(ctx context.Context) ([]models.Record, error) {
	query := `SELECT id, title, total_episodes, watched_episodes, type, status FROM watch_list`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.Record
	for rows.Next() {
		var record models.Record
		err := rows.Scan(&record.ID,
			&record.Title, &record.TotalEpisodes,
			&record.WatchedEpisodes, &record.Type,
			&record.Status)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *RecordRepository) UpdateRecord(ctx context.Context, record *models.Record) error {
	query := `
	UPDATE watch_list
	SET title = ?, total_episodes = ?, watched_episodes = ?,
	type = ?, status = ? WHERE id = ?
	`

	result, err := r.db.ExecContext(ctx, query,
		record.Title, record.TotalEpisodes,
		record.WatchedEpisodes, record.Type,
		record.Status, record.ID)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no record found with ID %d", record.ID)
	}

	return nil
}

func (r *RecordRepository) DeleteRecord(ctx context.Context, id int) error {
	query := `DELETE FROM watch_list WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no record found with ID %d", id)
	}

	return nil
}
