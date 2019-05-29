package db

import (
	"encoding/json"
	"fmt"
	"os"

	// import sql
	_ "database/sql"
	// import mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/juju/loggo"
	"github.com/xhan7279/approvalAssistant/config"
)

// DB is the shared database connection
var (
	DB     *gorm.DB
	logger = loggo.GetLogger("DB")
)

// InitDb to open connection to Mysql
func InitDb() error {
	fpath := os.Getenv("CONFIG_FILE")
	logger.Infof("Reading config file =", fpath)
	fmt.Println("config=", fpath)

	connstr, err := getConfig(fpath)
	if err != nil {
		logger.Errorf("Failed to get config", err)
		return err
	}
	DB, err := gorm.Open("mysql", connstr)

	if err != nil {
		logger.Errorf("Failed to open connection", err)
		return err
	}
	_ = DB
	return nil
}

// getConfig return a struct Configuration with server configs
func getConfig(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	decoder := json.NewDecoder(file)
	configuration := config.Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		return "", err
	}
	return configuration.ConnectionString(), nil
}
