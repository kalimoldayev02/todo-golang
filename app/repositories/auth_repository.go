package repositories

import (
	"database/sql"
	"fmt"
	"strings"
	"todo-golang/app/dto"
	"todo-golang/app/entities"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(signUpDto dto.SignUpDto) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, signUpDto.FirstName, signUpDto.LastName, strings.ToLower(signUpDto.Email), signUpDto.Password)
	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (r *AuthRepository) GetUserByCredentials(email, password string) (entities.User, error) {
	var userEntity entities.User
	query := fmt.Sprintf("select id, first_name, last_name, email, password, TO_CHAR(created_at, 'DD-MM-YYYY') from %s WHERE LOWER(email)=$1 AND password=$2", usersTable)

	err := r.db.QueryRow(query, strings.ToLower(email), password).Scan(
		&userEntity.Id,
		&userEntity.FirstName,
		&userEntity.LastName,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return entities.User{}, fmt.Errorf("user not found")
		}
		return entities.User{}, err
	}

	return userEntity, nil
}
