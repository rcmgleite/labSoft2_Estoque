package database

import (
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	//Blank import needed to call init from go-sqlite3()
	_ "github.com/mattn/go-sqlite3"
)

const (
	databaseType     = "sqlite3"
	databaseFile     = "./estoque.db"
	databaseTestFile = "./estoque_test.db"
)

//GetDatabase ...
func GetDatabase() *gorm.DB {
	testing, _ := strconv.ParseBool(os.Getenv("TEST"))
	var dbPath string
	if testing {
		dbPath = databaseTestFile
	} else {
		dbPath = databaseFile
	}
	db, err := gorm.Open(databaseType, dbPath)
	if err != nil {
		panic(err)
	}
	return &db
}
