package repositories

import (
	"database/sql"
	"todo-golang/app/dto"
)

const (
	usersTable = "users"
)

type Auth interface {
	SignUp(signUpDto dto.SignUpDto) (int, error)
}

type Respository struct {
	Auth
}

func NewRepository(db *sql.DB) *Respository {
	return &Respository{
		Auth: NewAuthRepository(db),
	}
}
