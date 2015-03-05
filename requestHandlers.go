package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var store = NewProductStore()

// var orderStore = NewOrderStore(1)
var uOrder = NewOrder()

// defaultHandler Just redirect the incomming default "/" request to index
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

// productHandler - handles all http methods for the "/product"
// POST: add new product
// GET: retreive product
// DELETE: delete product
// PUT: update product
func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//Gambi.. enquanto não usa JQUERY não da pra fazer nada de diferente de GET e POST
		//http://stackoverflow.com/questions/1856996/doing-a-http-put-from-a-browser
		if r.URL.Query().Get("delete") != "" {
			//TODO
		} else {
			t, _ := template.ParseFiles("product.html")
			t.Execute(w, store.Products)
		}

		return

	case "POST":
		name := r.FormValue("name")
		description := r.FormValue("description")
		currQuantity, errCurr := strconv.ParseInt(r.FormValue("currQuantity"), 0, 64)
		minQuantity, errMin := strconv.ParseInt(r.FormValue("minQuantity"), 0, 64)

		if errCurr != nil || errMin != nil || description == "" {
			http.Redirect(w, r, "/product", http.StatusFound)
			return
		}

		p := Product{name, description, currQuantity, minQuantity}
		store.AddProduct(p)

		if p.needRefill() {
			uOrder.addItem(p)
		}

		http.Redirect(w, r, "/product", http.StatusFound)
		return

	case "PUT":
		fmt.Println("Put method used!")
		return

	case "DELETE":
		fmt.Println("Delete method used!")
		return
	}
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		t, _ := template.ParseFiles("order.html")
		t.Execute(w, uOrder)
		return

	case "POST":
		uOrder.approve()
		http.Redirect(w, r, "/order", http.StatusFound)

		return

	case "PUT":
		fmt.Println(r.Method)
		return

	case "DELETE":
		fmt.Println("Delete method used!")
		return
	}
}
