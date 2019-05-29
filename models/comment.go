// Package models contains utility functions for database mapping.
package models

import (
	"time"
)

// CommentModel represents one distinct project with unique id
type CommentModel struct {
	cID     int64  `gorm:"primary_key"`
	comment string `gorm:"type:varchar(100)"`
	created int64
	author  int64 `gorm:"ForeignKey:user_id"`
}

// CreateComment creates a comment made by user
func CreateComment(comment string, uID int64) CommentModel {
	return CommentModel{
		comment: comment,
		created: time.Now().Unix(),
		author:  uID,
	}
}
