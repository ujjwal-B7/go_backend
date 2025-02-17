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

func (r *Repository) GetAllFutsals() (*[]model.Futsal, error) {

	var data []model.Futsal

	results, err := r.DB.Query("SELECT * FROM futsal")

	if err != nil {
		return nil, err
	}

	defer results.Close()

	for results.Next() {

		var result model.Futsal

		err := results.Scan(&result.ID, &result.Name, &result.Phone, &result.Price, &result.Location)

		if err != nil {
			panic(err.Error())
		}

		data = append(data, result)
	}

	return &data, nil
}

// update for all the fields at once, if any field is not provide in update case it becomes empty in database
func (r *Repository) UpdateFutsal(id uint64, futsal model.Futsal) (*model.Futsal, error) {

	statement, err := r.DB.Prepare("UPDATE futsal SET name=?,phone=?,price=?,location=? WHERE id=?")

	if err != nil {
		return nil, err
	}

	defer statement.Close()
	_, err = statement.Exec(futsal.Name, futsal.Phone, futsal.Price, futsal.Location, id)

	if err != nil {
		return nil, err
	}

	updatedData := &model.Futsal{}
	err = r.DB.QueryRow("SELECT id,name,phone,price,location FROM futsal WHERE id=?", id).Scan(
		&updatedData.ID,
		&updatedData.Name,
		&updatedData.Phone,
		&updatedData.Price,
		&updatedData.Location,
	)

	if err != nil {
		return nil, err
	}
	return updatedData, nil
}

func (r *Repository) UpdateFutsalFields(id uint64, futsal model.Futsal) (*model.Futsal, error) {

	query := "UPDATE futsal SET "

	setClauses := []string{}
	var args []interface{}

	if futsal.Name != "" {
		setClauses = append(setClauses, "name=?")
		args = append(args, futsal.Name)
	}
	if futsal.Price != "" {
		setClauses = append(setClauses, "price=?")
		args = append(args, futsal.Price)
	}
	if futsal.Phone != "" {
		setClauses = append(setClauses, "phone=?")
		args = append(args, futsal.Phone)
	}

	if futsal.Location != "" {
		setClauses = append(setClauses, "location=?")
		args = append(args, futsal.Location)
	}

	for idx, clause := range setClauses {
		if idx == len(setClauses)-1 {
			query += clause
			break
		}
		query += clause + ","
	}

	query += " WHERE id=?"

	statement, err := r.DB.Prepare(query)
	args = append(args, id)

	if err != nil {
		return nil, err
	}

	defer statement.Close()
	_, err = statement.Exec(args...)

	if err != nil {
		return nil, err
	}

	updatedData := &model.Futsal{}
	err = r.DB.QueryRow("SELECT id,name,phone,price,location FROM futsal WHERE id=?", id).Scan(
		&updatedData.ID,
		&updatedData.Name,
		&updatedData.Phone,
		&updatedData.Price,
		&updatedData.Location,
	)

	if err != nil {
		return nil, err
	}
	return updatedData, nil
}

func (r *Repository) DeleteFutsal(id uint64) error {

	statement, err := r.DB.Prepare("DELETE FROM futsal where id=?")

	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
