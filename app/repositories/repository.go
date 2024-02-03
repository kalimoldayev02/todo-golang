package repositories

import (
	"database/sql"
	"todo-golang/app/dto"
	"todo-golang/app/entities"
)

const (
	usersTable = "users"
)

type Auth interface {
	CreateUser(signUpDto dto.SignUpDto) (int, error)
	GetUserByCredentials(email, passwors string) (entities.User, error)
}

type Respository struct {
	Auth
}

func NewRepository(db *sql.DB) *Respository {
	return &Respository{
		Auth: NewAuthRepository(db),
	}
}
