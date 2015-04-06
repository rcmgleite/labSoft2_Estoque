package main

import (
	"net/http"

	"github.com/rcmgleite/labSoft2_Estoque/models"
	"github.com/rcmgleite/labSoft2_Estoque/requestDecoder"
)

//POSTQueryProductHandler ...
func POSTQueryProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	decoder := requestDecoder.NewDecoder()
	err := decoder.DecodeReqBody(&p, r.Body)

	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	products, err := p.Retreive()
	rj := NewResponseJSON(products, err)
	writeBack(w, r, rj)

}

// GETProductHandler ...
func GETProductHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	var p models.Product

	decoder := requestDecoder.NewDecoder()
	err := decoder.DecodeURLValues(&p, queryString)

	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	products, err := p.Retreive()
	rj := NewResponseJSON(products, err)
	writeBack(w, r, rj)
}

// FIXME - make database insertions on the same transaction
func addProductToOrder(p models.Product) {
	var order models.Order
	err := order.GetOpenOrder()
	if err != nil && err.Error() == "record not found" {
		order.Save()
		order.AddProduct(p)
	} else {
		order.AddProduct(p)
	}
}

// POSTProductHandler ...
func POSTProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	decoder := requestDecoder.NewDecoder()
	err := decoder.DecodeReqBody(&p, r.Body)
	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	err = p.Save()
	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	if p.NeedRefill() {
		addProductToOrder(p)
	}
	rj := NewResponseJSON("Product successfully saved", err)
	writeBack(w, r, rj)
}

// PUTProductHandler ...
func PUTProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	decoder := requestDecoder.NewDecoder()
	err := decoder.DecodeReqBody(&p, r.Body)

	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	err = p.Update()

	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	if p.NeedRefill() {
		addProductToOrder(p)
	}

	rj := NewResponseJSON("Product updated successfully", err)
	writeBack(w, r, rj)
}

// DELETEProductHandler ...
func DELETEProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	decoder := requestDecoder.NewDecoder()
	err := decoder.DecodeReqBody(&p, r.Body)

	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	err = p.Delete()

	rj := NewResponseJSON("Product deleted successully", err)
	writeBack(w, r, rj)
}

// GETOrderHandler ...
func GETOrderHandler(w http.ResponseWriter, r *http.Request) {
	order := models.Order{}
	order.GetOpenOrder()
	writeBack(w, r, order)
}

// PUTOrderHandler ...
func PUTOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	decoder := requestDecoder.NewDecoder()
	err := decoder.DecodeReqBody(&order, r.Body)

	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	err = order.Update()

	rj := NewResponseJSON("Order updated successfully", err)
	writeBack(w, r, rj)
}

// DELETEOrderHandler ...
func DELETEOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	decoder := requestDecoder.NewDecoder()
	err := decoder.DecodeReqBody(&order, r.Body)

	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	err = order.Delete()

	rj := NewResponseJSON("Order deleted successfully", err)
	writeBack(w, r, rj)
}
