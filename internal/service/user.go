package service

import (
	"todo-golang/internal/model"
	"todo-golang/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return nil, nil
}
