package main

import (
	"net/http"
	"strconv"
	"text/template"
)

var database = NewDatabase()
var mPedido = NewPedido()

func produtoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t, _ := template.ParseFiles("produto.html")
		t.Execute(w, database.Produtos)

		return

	case "POST":
		name := r.FormValue("name")
		currQuantity, errCurr := strconv.ParseInt(r.FormValue("currQuantity"), 0, 64)
		minQuantity, errMin := strconv.ParseInt(r.FormValue("minQuantity"), 0, 64)

		if errCurr != nil || errMin != nil || name == "" {
			http.Redirect(w, r, "/produto", http.StatusFound)
			return
		}

		p := Produto{name, currQuantity, minQuantity}
		database.AddProduto(p)

		if p.precisaRepor() {
			mPedido.addItem(p)
		}

		http.Redirect(w, r, "/produto", http.StatusFound)
		return
	}
}

func pedidoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		t, _ := template.ParseFiles("pedido.html")
		t.Execute(w, mPedido)
		return

	case "POST":
		mPedido.approve()
		http.Redirect(w, r, "/pedido", http.StatusFound)

		return
	}
}
