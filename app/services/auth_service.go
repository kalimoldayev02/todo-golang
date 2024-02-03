package services

import (
	"todo-golang/app/dto"
	"todo-golang/app/repositories"
	"todo-golang/pkg/utils"
)

type AuthService struct {
	repository repositories.Auth
}

func NewAuthSerive(r repositories.Auth) *AuthService {
	return &AuthService{
		repository: r,
	}
}

func (s *AuthService) SignUp(signUpDto dto.SignUpDto) (int, error) {
	signUpDto.Password = utils.GeneratePasswordHash(signUpDto.Password)

	return s.repository.CreateUser(signUpDto)
}

func (s *AuthService) SignIn(signInDto dto.SignInDto) (string, error) {
	signInDto.Password = utils.GeneratePasswordHash(signInDto.Password)

	userEntity, err := s.repository.GetUserByCredentials(signInDto.Email, signInDto.Password)
	if err != nil {
		return "", err
	}

	return userEntity.Password, nil
}
