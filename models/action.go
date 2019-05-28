// Package models contains utility functions for database mapping.
package models

// ActionModel represents an event an user can perform
type ActionModel struct {
	aID  int64  `db:"a_id, primarykey, autoincrement"`
	name string `db:", size:20"`
}
