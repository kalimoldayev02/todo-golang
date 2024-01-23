package repository

import (
	"database/sql"
	"todo-golang/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRespository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	users := []model.User{}
	return users, nil
}
