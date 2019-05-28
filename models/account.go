// Package models contains utility functions for database mapping
package models

// Account represents one unique user
type Account struct {
	uID      int64  `db:"usr_id, primarykey, autoincrement"`
	username string `db:", size:30"`
	password string `db:", size:20"`
	created  int64
}

// CreateAccount
func CreateAccount(username, password string) Account {
	return Account{
		username: username,
		password: password,
	}
}
