// Package models contains utility functions for database mapping.
package models

import (
	"time"
)

// CommentModel represents one distinct project with unique id
type CommentModel struct {
	cID     int64  `db:"c_id, primarykey, autoincrement"`
	comment string `db:", size:100"`
	created int64
	author  int64 `db:", foreignkey"`
}

// CreateComment creates a comment made by user
func CreateComment(comment string, uID int64) CommentModel {
	return CommentModel{
		comment: comment,
		created: time.Now().Unix(),
		author:  uID,
	}
}
