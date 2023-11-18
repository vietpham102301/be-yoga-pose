package main

import (
	"net/http"
	"yoga-pose-backend/database"
	"yoga-pose-backend/routes"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := routes.SetupRoutes(db)
	http.ListenAndServe(":8080", r)
}
