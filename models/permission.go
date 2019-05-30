package models

import "time"

// PermissionModel represents one distinct project with unique id
type PermissionModel struct {
	pID         int64 `gorm:"primary_key"`
	actionID    int64
	userID      int64
	projectID   int64
	grantedtime int64
}

// CreatePermission creates a project structure
func CreatePermission(actionID, userID, projectID, grantedtime int64) PermissionModel {
	return PermissionModel{
		actionID:    actionID,
		userID:      userID,
		projectID:   projectID,
		grantedtime: time.Now().Unix(),
	}
}
