package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

//DbFactory struct - will be a Singleton factory to retreive the correct db
type DbFactory struct {
	dbType string
	dbs    map[string]*gorm.DB
}

var instance *DbFactory

//GetDbFactoryInstance = Singleton constructor
func GetDbFactoryInstance(dbType string) *DbFactory {
	var err error
	if instance == nil || instance.dbType != dbType {
		instance = nil
		instance = new(DbFactory)
		instance.dbs = make(map[string]*gorm.DB)
		instance.dbType = dbType
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return instance
}

//GetDatabase ...
func (dbF *DbFactory) GetDatabase(dbPath string) *gorm.DB {
	if dbF.dbs[dbPath] != nil {
		return dbF.dbs[dbPath]
	}

	db, err := gorm.Open(dbF.dbType, dbPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	dbF.dbs[dbPath] = &db
	return &db
}
