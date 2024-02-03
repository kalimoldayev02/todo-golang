package services

import (
	"todo-golang/app/dto"
	"todo-golang/app/repositories"
	"todo-golang/pkg/utils"
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
	signUpDto.Password = utils.GeneratePasswordHash(signUpDto.Password)

	return a.repository.SignUp(signUpDto)
}
