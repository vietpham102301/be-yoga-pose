package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"yoga-pose-backend/database"
	"yoga-pose-backend/python"
	"yoga-pose-backend/routes"
)

func main() {
	python.GetABSPath()
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}(db)

	r := routes.SetupRoutes(db)
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
