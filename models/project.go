// Package models contains utility functions for database mapping.
package models

import (
	"github.com/jinzhu/gorm"
)

// ProjectModel represents one distinct project with unique id
type ProjectModel struct {
	gorm.Model
	name    string `gorm:"type:varchar(100); not null"`
	duedate Date
	userID  int64 `gorm:"not null"`
}

// CreateProject creates a project structure
func CreateProject(name string, duedate Date, userID int64) ProjectModel {
	return ProjectModel{
		name:    name,
		duedate: duedate,
		userID:  userID,
	}
}
