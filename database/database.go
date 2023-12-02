package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"yoga-pose-backend/config"
)

func InitDB() (*sql.DB, error) {
	time.Sleep(10 * time.Second)

	db, err := sql.Open("mysql", config.DataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
