package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

//DatabaseFactory struct - will be a Singleton factory to retreive the correct db
type DatabaseFactory struct {
	dbType string
}

var instance *DatabaseFactory

//getDbFactoryInstance = Singleton constructor
func getDbFactoryInstance(dbType string) *DatabaseFactory {
	var err error
	if instance == nil || instance.dbType != dbType {
		instance = nil
		instance = new(DatabaseFactory)
		instance.dbType = dbType
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return instance
}

func (dbF *DatabaseFactory) getDataBase(dbPath string) *gorm.DB {
	db, err := gorm.Open(dbF.dbType, dbPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &db
}
