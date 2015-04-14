package main

import (
	"net/http"

	"github.com/rcmgleite/labSoft2_Estoque/decoder"
	"github.com/rcmgleite/labSoft2_Estoque/models"
)

//POSTQueryProductHandler ...
func POSTQueryProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	decoder := decoder.NewDecoder()
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

	decoder := decoder.NewDecoder()
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

// POSTProductHandler ...
func POSTProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	decoder := decoder.NewDecoder()
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

	rj := NewResponseJSON("Product successfully saved", err)
	writeBack(w, r, rj)
}

// PUTProductHandler ...
func PUTProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	decoder := decoder.NewDecoder()
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

	rj := NewResponseJSON("Product updated successfully", err)
	writeBack(w, r, rj)
}

// DELETEProductHandler ...
func DELETEProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	decoder := decoder.NewDecoder()
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
