package repository

import (
	"database/sql"
	"todo-golang/internal/model"
)

type User interface {
	GetUsers() ([]model.User, error)
}

type Resporitory struct {
	User
}

func NewRespository(db *sql.DB) *Resporitory {
	return &Resporitory{
		User: NewUserRespository(db),
	}
}
