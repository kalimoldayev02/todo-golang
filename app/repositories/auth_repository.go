package repositories

import (
	"database/sql"
	"fmt"
	"todo-golang/app/dto"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) SignUp(signUpDto dto.SignUpDto) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, signUpDto.FirstName, signUpDto.LastName, signUpDto.Email, signUpDto.Password)
	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}
