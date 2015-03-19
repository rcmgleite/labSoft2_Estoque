package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

//Config variables
var dbType = "sqlite3"

//DatabaseFactory struct - will be a Singleton factory to retreive the correct db
type DatabaseFactory struct {
	db gorm.DB
}

var instance *DatabaseFactory

//NewDatabase = Singleton constructor
func getDbFactoryInstance(dbPath string) *DatabaseFactory {
	var err error
	if instance == nil {
		instance = new(DatabaseFactory)
		instance.db, err = gorm.Open(dbType, dbPath)
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return instance
}

func (dbF *DatabaseFactory) getDataBase() *gorm.DB {
	return &dbF.db
}
