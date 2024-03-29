package dto

type SignUpDto struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type SignInDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
