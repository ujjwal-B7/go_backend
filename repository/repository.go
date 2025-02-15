package repository

import (
	"backend/model"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) model.FutsalRepositoryInterface {
	return &Repository{DB: db}
}

func (r *Repository) SaveFutsal(futsal model.Futsal) error {

	statement, err := r.DB.Prepare("INSERT INTO futsal(name,phone,price,location) VALUES (?,?,?,?)")

	if err != nil {
		return err
	}

	_, err = statement.Exec(futsal.Name, futsal.Phone, futsal.Price, futsal.Location)

	if err != nil {
		return err
	}

	return nil
}
