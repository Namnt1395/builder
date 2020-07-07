package mysql

import (
	"builder/demo_querybuilder/model/adapters"
	"fmt"
)

var database *adapters.MysqlAdapter

// OpenDatabase opens the database with the given options
func OpenDatabase() error {
	if database != nil {
		return fmt.Errorf("query: database already open - %s", database)
	}
	database = &adapters.MysqlAdapter{}
	if database == nil {
		return fmt.Errorf("query: database adapter not recognised")
	}
	return database.Open()
}

// CloseDatabase closes the database opened by OpenDatabase
func CloseDatabase() error {
	var err error
	if database != nil {
		err = database.Close()
		database = nil
	}
	return err
}
