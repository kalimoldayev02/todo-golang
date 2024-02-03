package services

import (
	"todo-golang/app/dto"
	"todo-golang/app/repositories"
)

type AuthService struct {
	repository repositories.Auth
}

func NewAuthSerive(repo repositories.Auth) *AuthService {
	return &AuthService{
		repository: repo,
	}
}

func (a *AuthService) SignUp(signUpDto dto.SignUpDto) (int, error) {
	return a.repository.SignUp(signUpDto)
}
