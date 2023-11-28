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

func (hr *HistoryPoseRepository) SaveHistoryLog(name string, path string, userID int, score int) error {
	query := "INSERT INTO history (pose_name, path, user_id, score, is_saved) VALUES (?, ?, ?, ?, ?)"
	_, err := hr.db.Exec(query, name, path, userID, score, 0)
	if err != nil {
		return err
	}
	return nil
}

func (hr *HistoryPoseRepository) GetHistoryLog(pageSize int, pageNum int, userID int) (*[]models.History, error) {
	query := "SELECT * FROM history WHERE user_id = ?  ORDER BY created_at DESC LIMIT ? OFFSET ?"
	rows, err := hr.db.Query(query, userID, pageSize, pageNum)
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
		err := rows.Scan(&h.ID, &h.UserID, &h.Name, &h.Score, &h.Path, &h.CreatedAt, &h.IsSaved)
		if err != nil {
			return nil, err
		}
		*history = append(*history, *h)
	}
	return history, nil
}

func (hr *HistoryPoseRepository) GetHistoryByID(id int) (*models.History, error) {
	query := "SELECT * FROM history WHERE id = ?"
	_, err := hr.db.Exec(query, id)

	if err != nil {
		return nil, err
	}
	history := &models.History{}
	err = hr.db.QueryRow(query, id).Scan(&history.ID, &history.UserID, &history.Name, &history.Score, &history.Path, &history.CreatedAt, &history.IsSaved)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (hr *HistoryPoseRepository) DeleteHistoryByID(id int, userID int) error {
	query := "DELETE FROM history WHERE id = ? AND user_id = ?"
	_, err := hr.db.Exec(query, id, userID)
	if err != nil {
		return err
	}
	return nil
}

func (hr *HistoryPoseRepository) SavedHistory(id int, userID int) error {
	query := "UPDATE history SET is_saved = true WHERE id = ? AND user_id = ?"
	_, err := hr.db.Exec(query, id, userID)
	if err != nil {
		return err
	}
	return nil
}
