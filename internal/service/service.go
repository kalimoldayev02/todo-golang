package service

import (
	"todo-golang/internal/model"
	"todo-golang/internal/repository"
)

type User interface {
	GetUsers() ([]model.User, error)
}

type Service struct {
	User
}

func NewService(repo *repository.Resporitory) *Service {
	return &Service{
		User: NewUserService(repo.User),
	}
}
