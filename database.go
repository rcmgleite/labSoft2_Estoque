package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

//Config variables
var dbPath = "./estoque.db"
var dbType = "sqlite3"

//Database struct - will be a Singleton
type Database struct {
	db gorm.DB
}

var instance *Database

//NewDatabase = Singleton constructor
func getDbInstance() *Database {
	var err error
	if instance == nil {
		instance = new(Database)
		instance.db, err = gorm.Open(dbType, dbPath)
		fmt.Println("First db instance created")
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("returning an already allocated instance after calling NewDao()")
	return instance
}
