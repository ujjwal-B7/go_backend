package controller

import (
	"backend/model"
	"backend/repository"
	"backend/usecase"
	"database/sql"
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	DB      *sql.DB
	Router  *mux.Router
	UseCase model.FutsalUsecaseInterface
}

func NewServer(db *sql.DB) *Server {
	repo := repository.NewRepository(db)
	useCase := usecase.NewUsecase(repo)
	server := &Server{
		DB:      db,
		Router:  mux.NewRouter(),
		UseCase: useCase,
	}

	server.initializeRoutes()

	logrus.Info("Server started...")
	return server
}

func ConnectDB() *sql.DB {

	dbUrl := fmt.Sprintf("%s:%s@tcp(%v:%v)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", dbUrl)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func MigrateDB(db *sql.DB) error {

	// create players
	createTableQuery := `
  		CREATE TABLE IF NOT EXISTS teammate(
  		id INT AUTO_INCREMENT PRIMARY KEY,
  		name VARCHAR(255),
  		email VARCHAR(255),
  		location VARCHAR(255),
  		phone VARCHAR(20),
  		playing_position VARCHAR(50),
  		age VARCHAR(20)
        )
    `

	_, err := db.Exec(createTableQuery)

	if err != nil {
		fmt.Print("Failed to create teammate table", err)
		return err
	}

	createFutsal := `
	CREATE TABLE IF NOT EXISTS futsal(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		phone VARCHAR(20),
		price VARCHAR(255),
		location VARCHAR(255)
        )
		`
	_, err = db.Exec(createFutsal)
	if err != nil {
		fmt.Print("Failed to create futsal table",err)
		return err
	}

	return nil
}
