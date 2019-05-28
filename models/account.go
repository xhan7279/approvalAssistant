// Package models contains utility functions for database mapping
package models

// AccountModel represents one unique user
type AccountModel struct {
	uID      int64  `db:"usr_id, primarykey, autoincrement"`
	username string `db:", size:30"`
	password string `db:", size:20"`
	created  int64
}

// CreateAccount creates an account based on username/password combination
func CreateAccount(username, password string) AccountModel {
	return AccountModel{
		username: username,
		password: password,
	}
}
