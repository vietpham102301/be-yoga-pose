package repository

import (
	"database/sql"
	"yoga-pose-backend/models"
)

type SavePoseRepository struct {
	db *sql.DB
}

func NewSavePoseRepository(db *sql.DB) *SavePoseRepository {
	return &SavePoseRepository{db}
}

func (sr *SavePoseRepository) SavePose(userID int, name string, path string) error {
	query := "INSERT INTO saved_pose (user_id, name, path) VALUES (?, ?, ?)"
	_, err := sr.db.Exec(query, userID, name, path)
	if err != nil {
		return err
	}
	return nil
}

func (sr *SavePoseRepository) GetSavedPoses(pageSize int, pageNum int, userID int) (*[]models.SavedPose, error) {
	query := "SELECT * FROM saved_pose LIMIT ? OFFSET ? WHERE user_id = ?"
	rows, err := sr.db.Query(query, pageSize, pageNum, userID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	savedPoses := &[]models.SavedPose{}
	for rows.Next() {
		s := &models.SavedPose{}
		err := rows.Scan(&s.ID, &s.UserID, &s.Name, &s.Path, &s.CreatedAt)
		if err != nil {
			return nil, err
		}
		*savedPoses = append(*savedPoses, *s)
	}
	return savedPoses, nil
}
