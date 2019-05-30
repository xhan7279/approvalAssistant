// Package models contains utility functions for database mapping
package models

import (
	"github.com/jinzhu/gorm"
)

// AccountModel represents one unique user
type AccountModel struct {
	gorm.Model
	username string `gorm:"type:varchar(30);not null"`
	password string `gorm:"type:varchar(30);not null"`
	email    string `gorm:"type:varchar(100);not null"`
}

// CreateAccount creates an account based on username/password combination
func CreateAccount(username, password, email string) AccountModel {
	return AccountModel{
		username: username,
		password: password,
		email:    email,
	}
}
