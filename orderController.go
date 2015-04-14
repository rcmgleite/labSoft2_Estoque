package main

import (
	"net/http"

	"github.com/rcmgleite/labSoft2_Estoque/decoder"
	"github.com/rcmgleite/labSoft2_Estoque/models"
)

// GETOrderHandler ...
func GETOrderHandler(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	o.GetOpenOrder()
	rj := NewResponseJSON(o, nil)
	writeBack(w, r, rj)
}

// PUTOrderHandler ...
func PUTOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	decoder := decoder.NewDecoder()
	err := decoder.DecodeReqBody(&order, r.Body)

	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	err = order.Update()
	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	rj := NewResponseJSON("Order updated successfully", err)
	writeBack(w, r, rj)
}

// DELETEOrderHandler ...
func DELETEOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	decoder := decoder.NewDecoder()
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
