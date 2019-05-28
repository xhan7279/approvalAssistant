// Package models contains utility functions for database mapping.
package models

// File represents an attachment to a project
type FileModel struct {
	fID  int64  `db:"f_id, primarykey, autoincrement"`
	name string `db:", size:100"`
	pid  int64  `db:"p_id, foreignkey"`
}
