package main

import (
	"backend/controller"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading env %v", err)
	}

	db := controller.ConnectDB()

	server := controller.NewServer(db)
	controller.MigrateDB(db)

	log.Fatal(http.ListenAndServe(":8000", server.Router))
}
