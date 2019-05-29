// Package models contains utility functions for database mapping.
package models

// FileModel represents an attachment to a project
type FileModel struct {
	fID  int64  `gorm:"primary_key"`
	name string `gorm:"type:varchar(100)"`
	pid  int64  `gorm:"p_id,foreignkey"`
}
