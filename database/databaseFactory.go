package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	//Blank import needed
	_ "github.com/mattn/go-sqlite3"
)

//DbFactory struct - will be a Singleton factory to retreive the correct db
type DbFactory struct {
	dbType string
	dbs    map[string]*Database
}

// Database ...
type Database struct {
	gorm.DB
	transaction *gorm.DB
	commitCalls int
	mutex       sync.RWMutex
}

var instance *DbFactory

//GetDbFactoryInstance = Singleton constructor
func GetDbFactoryInstance(dbType string) *DbFactory {
	var err error
	if instance == nil || instance.dbType != dbType {
		instance = nil
		instance = new(DbFactory)
		instance.dbs = make(map[string]*Database)
		instance.dbType = dbType
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return instance
}

//GetDatabase ...
func (dbF *DbFactory) GetDatabase(dbPath string) *Database {
	if dbF.dbs[dbPath] != nil {
		return dbF.dbs[dbPath]
	}

	db, err := gorm.Open(dbF.dbType, dbPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	dbF.dbs[dbPath] = &Database{db, nil, 0, sync.RWMutex{}}
	return dbF.dbs[dbPath]
}

// GetTransaction ... Used to avoid nested transactions
func (db *Database) GetTransaction() *gorm.DB {
	db.mutex.Lock()
	db.commitCalls++
	db.mutex.Unlock()
	if db.transaction != nil {
		return db.transaction
	}
	db.transaction = db.Begin()
	return db.transaction
}

// DoCommit ...
func (db *Database) DoCommit() {
	db.mutex.Lock()
	db.commitCalls--
	db.mutex.Unlock()

	if db.commitCalls > 0 {
		return
	} else if db.commitCalls == 0 {
		db.transaction.Commit()

		//Delete already commited transaction
		db.transaction = nil
		return
	} else {
		log.Fatal("Should never get here - function DoCommit on Database struct")
	}

}

// DoRollback ...
func (db *Database) DoRollback() {
	db.transaction.Rollback()
	db.commitCalls = 0

	//Delete rollbacked transaction
	db.transaction = nil
}
