package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"yoga-pose-backend/config"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.GetMySQLURL())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
