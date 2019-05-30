// Package models contains utility functions for database mapping.
package models

import (
	"github.com/jinzhu/gorm"
)

// FileModel represents an attachment to a project
type FileModel struct {
	gorm.Model
	name      string `gorm:"type:varchar(100)"`
	projectID int64  `gorm:"foreignkey"`
	userID    int64  `gorm:"foreignkey"`
}

// CreateFile creates a project structure
func CreateFile(name string, projectID int64, userID int64) FileModel {
	return FileModel{
		name:      name,
		projectID: projectID,
		userID:    userID,
	}
}
