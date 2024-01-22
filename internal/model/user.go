package model

import "time"

type User struct {
	Id        uint
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	EditedAt  time.Time
}
