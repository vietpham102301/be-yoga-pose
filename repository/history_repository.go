package repository

import (
	"database/sql"
	"yoga-pose-backend/models"
)

type HistoryPoseRepository struct {
	db *sql.DB
}

func NewHistoryRepository(db *sql.DB) *HistoryPoseRepository {
	return &HistoryPoseRepository{db}
}

func (hr *HistoryPoseRepository) SaveHistoryLog(name string, path string, userID int) error {
	query := "INSERT INTO history (pose_name, path, user_id) VALUES (?, ?, ?)"
	_, err := hr.db.Exec(query, name, path, userID)
	if err != nil {
		return err
	}
	return nil
}

func (hr *HistoryPoseRepository) GetHistoryLog(pageSize int, pageNum int, userID int) (*[]models.History, error) {
	query := "SELECT * FROM history LIMIT ? OFFSET ? WHERE user_id = ?"
	rows, err := hr.db.Query(query, pageSize, pageNum, userID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	history := &[]models.History{}
	for rows.Next() {
		h := &models.History{}
		err := rows.Scan(&h.ID, &h.Name, &h.Path, &h.CreatedAt)
		if err != nil {
			return nil, err
		}
		*history = append(*history, *h)
	}
	return history, nil
}
