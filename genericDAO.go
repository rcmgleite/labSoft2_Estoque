package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//TODO refactor to be generic

//GenericDAO struct
type GenericDAO struct {
}

func test() {
	db, err := sql.Open("sqlite3", "./estoque.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;`

	//TODO verificar qual o primeiro parâmetro do retorno de db.Exec()
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}

}

func (dao *GenericDAO) save(p *Product) {
	//TODO
}

func (dao *GenericDAO) update(p *Product) {
	//TODO
}

func (dao *GenericDAO) delete(p *Product) {
	//TODO
}

func (dao *GenericDAO) retrive(p *Product) {
	//TODO
}
