// Package models contains utility functions for database mapping.
package models

import (
	"time"
)

type Timestamp time.Time

// Project represents one distinct project with unique id
type Project struct {
	pID     int64  `db:"p_id, primarykey, autoincrement"`
	name    string `db:", size:30"`
	created Timestamp
	due     Timestamp `db:", size:20"`
}

// CreateProject creates a project structure
func CreateProject(name string, due Timestamp) Project {
	return Project{
		name:    name,
		created: Timestamp(time.Now()),
		due:     due,
	}
}
