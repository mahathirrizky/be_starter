package user

import "time"

type UserTable struct {
	ID             int
	Name           string
	Email          string
	PasswordHash   string
	Role           string
	
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
