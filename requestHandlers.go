package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//defaultHandler Just redirect the incomming default "/" request to index
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html") //open and parse a template text file
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, nil)
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println("\n")
}
