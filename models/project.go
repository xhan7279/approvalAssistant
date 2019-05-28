// Package models contains utility functions for database mapping.
package models

import (
	"time"

	"github.com/fxtlabs/date"
)

// Date is a date
type Date date.Date

// ProjectModel represents one distinct project with unique id
type ProjectModel struct {
	pID     int64  `db:"p_id, primarykey, autoincrement"`
	name    string `db:", size:30"`
	created int64
	due     Date
}

// CreateProject creates a project structure
func CreateProject(name string, due Date) ProjectModel {
	return ProjectModel{
		name:    name,
		created: time.Now().Unix(),
		due:     due,
	}
}
