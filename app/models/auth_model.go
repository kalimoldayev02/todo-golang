package models

type SignUp struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required, email"`
	Password  string `json:"password" validate:"required, gte=8"`
}

type SignIn struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
