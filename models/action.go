// Package models contains utility functions for database mapping.
package models

import (
	"github.com/jinzhu/gorm"
)

// ActionModel represents an event an user can perform
type ActionModel struct {
	gorm.Model
	name string
}
