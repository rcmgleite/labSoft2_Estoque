package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//AddForm constant
const AddForm = `<div>
<form method="POST" action="/product">
Descrição: <input type="text" name="description"><br>
Quantidade atual: <input type="text" name="currQuantity"><br>
Quantidade mínima: <input type="text" name="minQuantity"><br>
<input type="submit" value="Adicionar">
</form></div>
`

//DefaultForm constant
const DefaultForm = `<div>
<form method="GET" action="/product">
<input type="submit" value="Produtos">
</form></div>
<div>
<form method="GET" action="/order">
<input type="submit" value="Pedidos de Compra">
</form></div>
`

//AproveFrom constant
const AproveForm = `<div>
<form method="POST" action="/order">
<input type="text" name="ignore" style="visibility:hidden"><br>
<input type="submit" value="Aprovar">
</form></div>
`

var store = NewProductStore(10)

// var orderStore = NewOrderStore(1)
var uOrder = NewOrder(10)

// defaultHandler Just redirect the incomming default "/" request to index
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, DefaultForm)
}

// productHandler - handles all http methods for the "/product"
// POST: add new product
// GET: retreive product
// DELETE: delete product
// PUT: update product
func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Query().Get("description") != "" {
			//TODO FALTA ATUALIZAR O NÚMERO MÍNIMO A QUALQUER MOMENTO
		} else {
			fmt.Fprintf(w, `<div><h2>Produtos cadastrados no Estoque: </h2></div>`)
			for i := 0; i < store.Size; i++ {
				fmt.Fprintf(w, "Descrição: %s, quantidade atual: %d, quantidade mínima: %d <br>", store.Products[i].description, store.Products[i].currQuantity, store.Products[i].minQuantity)

			}
			fmt.Fprintf(w, `<div><br><h2>Adicionar novo produto: </h2></div>`)
			fmt.Fprintf(w, AddForm)
		}

		return

	case "POST":
		id := store.Size
		description := r.FormValue("description")
		currQuantity, errCurr := strconv.ParseInt(r.FormValue("currQuantity"), 0, 64)
		minQuantity, errMin := strconv.ParseInt(r.FormValue("minQuantity"), 0, 64)

		if errCurr != nil || errMin != nil || description == "" {
			http.Redirect(w, r, "/product", http.StatusFound)
			return
		}

		p := Product{id, description, currQuantity, minQuantity}
		store.AddProduct(&p)

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
		fmt.Println(r.Method)
		fmt.Fprintf(w, `<div><h2>Pedido de compra: </h2></div>`)

		for i := 0; i < uOrder.size; i++ {
			fmt.Fprintf(w, "Descrição: %s -> quantidade a ser comprada: %d<br>", uOrder.productList[i].description, uOrder.productList[i].minQuantity-uOrder.productList[i].currQuantity)
		}

		if uOrder.size != 0 {
			if uOrder.approved == true {
				fmt.Fprintf(w, "Situação: Aprovada")
			} else {
				fmt.Fprintf(w, "Situação: Esperando Aprovação")
			}
		}
		if uOrder.approved != true && uOrder.size != 0 {
			fmt.Fprintf(w, AproveForm)
		}
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
