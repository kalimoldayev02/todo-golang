package repository

import (
	"database/sql"
	"fmt"
	"todo-golang/internal/model"
)

const tableName = "users"

type UserRepository struct {
	db *sql.DB
}

func NewUserRespository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	users := []model.User{}
	query := fmt.Sprintf("SELECT id, first_name, email FROM %s", tableName)

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uint
		var fitstName string
		var email string

		if err := rows.Scan(&id, &fitstName, &email); err != nil {
			return nil, err
		}

		user := model.User{
			Id:        id,
			FirstName: fitstName,
			Email:     email,
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println(users)

	return users, nil
}
