package repository

import (
	"database/sql"
	"yoga-pose-backend/models"
)

type YogaPoseRepository struct {
	db *sql.DB
}

func NewYogaRepository(db *sql.DB) *YogaPoseRepository {
	return &YogaPoseRepository{db}
}

func (yg *YogaPoseRepository) GetPoseByName(name string) (*models.Yoga, error) {
	query := "SELECT * FROM yoga_poses WHERE name = ?"
	yoga := &models.Yoga{}
	err := yg.db.QueryRow(query, name).Scan(&yoga.ID, &yoga.Name, &yoga.Path, &yoga.CreatedAt, &yoga.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return yoga, nil
}
