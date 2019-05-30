// Package models contains utility functions for database mapping.
package models

import (
	"time"
)

// ProjectModel represents one distinct project with unique id
type ProjectModel struct {
	pID         int64 `gorm:"primary_key"`
	name        string
	createdtime int64
	duedate     Date
	userID      int64
}

// CreateProject creates a project structure
func CreateProject(name string, duedate Date, userID int64) ProjectModel {
	return ProjectModel{
		name:        name,
		createdtime: time.Now().Unix(),
		duedate:     duedate,
		userID:      userID,
	}
}
