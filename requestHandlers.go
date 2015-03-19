package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var dao = newGenericDAO()

// defaultHandler Just redirect the incomming default "/" request to index
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

// GETProductHandler ...
func GETProductHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("product.html")
	t.Execute(w, dao.Retreive())
}

// POSTProductHandler ...
func POSTProductHandler(w http.ResponseWriter, r *http.Request) {
	var p Product

	if BuildStructFromForm(r, &p) {
		dao.Save(&p)
		if p.needRefill() {
			fmt.Println("will need refill")
		}
	}
	http.Redirect(w, r, "/product", http.StatusFound)

}

// PUTProductHandler ... TODO - add validations
func PUTProductHandler(w http.ResponseWriter, r *http.Request) {
	idFromForm, _ := strconv.Atoi(r.FormValue("id"))
	var p Product
	if BuildStructFromForm(r, &p) {
		p.ID = idFromForm
		dao.Update(&p)
	}
}

// DELETEProductHandler ...
func DELETEProductHandler(w http.ResponseWriter, r *http.Request) {
	idFromForm, _ := strconv.Atoi(r.FormValue("id"))
	p := Product{ID: idFromForm}
	dao.Delete(&p)
}

// GETOrderHandler ...
func GETOrderHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	t, _ := template.ParseFiles("order.html")
	t.Execute(w, nil)
}

// POSTOrderHandler ...
func POSTOrderHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	// uOrder.approve()
	http.Redirect(w, r, "/order", http.StatusFound)
}

// PUTOrderHandler ...
func PUTOrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUTOrderHandler")
	//TODO
}

// DELETEOrderHandler ...
func DELETEOrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETEOrderHandler")
	//TODO
}
