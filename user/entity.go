package user

import "time"

type User struct {
	ID             int
	Name           string
	Username       string
	Phone          string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
