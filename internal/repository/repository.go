package repository

import (
	"todo-golang/internal/model"

	"gorm.io/gorm"
)

type User interface {
	GetUsers() ([]model.User, error)
}

type Resporitory struct {
	User
}

func NewRespository(db *gorm.DB) *Resporitory {
	return &Resporitory{
		User: NewUserRespository(db),
	}
}
