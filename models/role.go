// Package models contains utility functions for database mapping.
package models

// RoleModel represents a role linked with permissions
type RoleModel struct {
	rID  int64  `db:"r_id, primarykey, autoincrement"`
	name string `db:", size:20"`
}
