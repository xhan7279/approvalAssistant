package models

import (
	"github.com/jinzhu/gorm"
)

// PermissionModel represents one distinct project with unique id
type PermissionModel struct {
	gorm.Model
	actionID  int64 `gorm:"type:varchar(100); not null"`
	userID    int64 `gorm:"not null"`
	projectID int64 `gorm:"not null"`
}

// CreatePermission creates a project structure
func CreatePermission(actionID, userID, projectID, grantedtime int64) PermissionModel {
	return PermissionModel{
		actionID:  actionID,
		userID:    userID,
		projectID: projectID,
	}
}
