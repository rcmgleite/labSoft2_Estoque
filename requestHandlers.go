package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rcmgleite/labSoft2_Estoque/models"
	"github.com/rcmgleite/labSoft2_Estoque/requestDecoder"
)

var comprasIP = "http://192.168.1.130:8080"

func makeRequest(httpMethod string, url string, requestObj []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(requestObj))
	addHeaders(req, headers)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func addHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

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

// PUTOrderHandler ...
func PUTOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	decoder := requestDecoder.NewDecoder()
	err := decoder.DecodeReqBody(&order, r.Body)

	fmt.Println(order.Approved)
	fmt.Println(order.ID)

	if err != nil {
		rj := NewResponseJSON(nil, err)
		writeBack(w, r, rj)
		return
	}

	err = order.Update()

	order.GetByID(order.ID)

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
			resp, err := makeRequest("POST", comprasIP+"/order", bJSON, headers)
			fmt.Println(err)
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("Response", string(body))
		}
	}

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
