package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UserName string
	Email    string
	Password string
}

func NewUser(username string, email string, password string) *User {
	return &User{
		UserName: username,
		Email:    email,
		Password: password,
	}
}
