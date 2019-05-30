// Package models contains utility functions for database mapping.
package models

import (
	"github.com/jinzhu/gorm"
)

// CommentModel represents one distinct project with unique id
type CommentModel struct {
	gorm.Model
	comment string `gorm:"type:varchar(100); not null"`
	userID  int64  `gorm:"ForeignKey:user_id"`
}

// CreateComment creates a comment made by user
func CreateComment(comment string, uID int64) CommentModel {
	return CommentModel{
		comment: comment,
		userID:  uID,
	}
}
