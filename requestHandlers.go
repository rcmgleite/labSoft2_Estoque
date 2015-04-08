package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rcmgleite/labSoft2_Estoque/decoder"
	"github.com/rcmgleite/labSoft2_Estoque/models"
	"github.com/rcmgleite/labSoft2_Estoque/requestHelper"
)

var comprasIP = "http://192.168.1.130:8080"

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

// GETOrderHandler ...
func GETOrderHandler(w http.ResponseWriter, r *http.Request) {
	// order := models.Order{}
	// order.GetOpenOrder()
	// writeBack(w, r, order)
	var o models.Order
	o.GetOpenOrder()
	rj := NewResponseJSON(o, nil)
	writeBack(w, r, rj)
}

func getJSON(object interface{}) ([]byte, error) {
	if object != nil {
		return json.Marshal(object)
	}
	return nil, nil
}

// Function that sends the new order to compras module
func sendOrderTOCompras(order *models.Order) error {
	err := order.GetByID(order.ID)

	if err == nil {
		headers := make(map[string]string)
		headers["Content-Type"] = "application/json"

		pToSend := make([]models.ProductToSend, len(order.Products))
		for index, value := range order.Products {
			pToSend[index].ProductID = value.ID
			pToSend[index].Quantidade = value.MinQuantity - value.CurrQuantity
			pToSend[index].Valor = 10
		}

		orderToSend := &models.OrderToSend{Products: pToSend}
		bJSON, err := getJSON(orderToSend)
		if err == nil {
			resp, err := requestHelper.MakeRequest("POST", comprasIP+"/order", bJSON, headers)
			if resp != nil && err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println("Response", string(body))
			} else {
				fmt.Println(err)
			}
		}
	}
	return err
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

	err = sendOrderTOCompras(&order)

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
