package services

import (
	"todo-golang/app/dto"
	"todo-golang/app/repositories"
)

type Auth interface {
	SignUp(signUpDto dto.SignUpDto) (int, error)
	SignIn(signInDto dto.SignInDto) (string, error)
}

type Service struct {
	Auth
}

func NewService(repo *repositories.Respository) *Service {
	return &Service{
		Auth: NewAuthSerive(repo.Auth),
	}
}
