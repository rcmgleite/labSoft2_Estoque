package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/rcmgleite/labEngSoft_Estoque/models"
)

type responseMSG struct {
	Msg string
}

var productDAO = newProductDAO()
var orderDAO = newOrderDAO()

func writeBack(w http.ResponseWriter, r *http.Request, i interface{}) {
	ct := r.Header.Get("Content-Type")
	switch ct {
	case "application/json":
		bJSON, err := json.Marshal(i)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(bJSON)
		break

	case "application/xml":
		bXML, err := xml.Marshal(i)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(bXML)
		break

	}
}

func createResponseMsg(err error) responseMSG {
	var msg string
	if err != nil {
		msg = err.Error()
	} else {
		msg = "Sucess"
	}
	return responseMSG{Msg: msg}
}

func parseReqBody(r *http.Request, i interface{}) error {
	r.ParseForm()
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(i)
}

// GETProductHandler ...
func GETProductHandler(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	productDAO.Retreive(&products)
	writeBack(w, r, products)
}

// POSTProductHandler ...
func POSTProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	err := parseReqBody(r, &p)
	if err != nil {
		fmt.Println(err)
	} else {
		err = productDAO.Save(&p)
		if err == nil {
			if p.NeedRefill() {
				//TODO
			}
		}
	}

	rj := createResponseMsg(err)

	writeBack(w, r, rj)
}

// PUTProductHandler ...
func PUTProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	err := parseReqBody(r, &p)

	if err == nil {
		err = productDAO.Update(&p)
	}

	rj := createResponseMsg(err)
	writeBack(w, r, rj)
}

// DELETEProductHandler ...
func DELETEProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	err := parseReqBody(r, &p)

	if err == nil {
		err = productDAO.Delete(&p)
	}

	rj := createResponseMsg(err)
	writeBack(w, r, rj)
}

// GETOrderHandler ...
func GETOrderHandler(w http.ResponseWriter, r *http.Request) {
	//FIXME
	order := models.Order{ID: -1, Approved: false}
	orderDAO.Retreive(&order)
	writeBack(w, r, order)
}

// PUTOrderHandler ...
func PUTOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := parseReqBody(r, &order)

	if err == nil {
		err = orderDAO.Update(&order)
	}

	rj := createResponseMsg(err)
	writeBack(w, r, rj)
}

// DELETEOrderHandler ...
func DELETEOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := parseReqBody(r, &order)

	if err == nil {
		err = orderDAO.Delete(&order)
	}

	rj := createResponseMsg(err)

	writeBack(w, r, rj)
}
