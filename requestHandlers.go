package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/rcmgleite/labSoft2_Estoque/models"
)

type responseMSG struct {
	Msg string
}

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

func queryStringToStruct(query url.Values, _struct interface{}) {
	newInstance := reflect.ValueOf(_struct).Elem()

	for k, v := range query {
		field := newInstance.FieldByName(k)

		switch field.Kind() {
		case reflect.Int:
			intValue, err := strconv.ParseInt(v[0], 0, 64)
			if err == nil {
				fmt.Println(field)
				field.SetInt(intValue)
			}

		case reflect.String:
			field.SetString(v[0])
		}
	}
}

// GETProductHandler ...
func GETProductHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	var p models.Product
	queryStringToStruct(queryString, &p)

	products, err := p.Retreive()
	if err != nil {
		rj := createResponseMsg(err)
		writeBack(w, r, rj)
	} else {
		writeBack(w, r, products)
	}
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
	err := parseReqBody(r, &p)
	if err != nil {
		fmt.Println(err)
	} else {
		err = p.Save()
		if err == nil {
			if p.NeedRefill() {
				addProductToOrder(p)
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
		err = p.Update()
		if err == nil {
			if p.NeedRefill() {
				addProductToOrder(p)
			}
		}
	}

	rj := createResponseMsg(err)
	writeBack(w, r, rj)
}

// DELETEProductHandler ...
func DELETEProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	err := parseReqBody(r, &p)

	if err == nil {
		err = p.Delete()
	}

	rj := createResponseMsg(err)
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
	err := parseReqBody(r, &order)

	if err == nil {
		err = order.Update()
	}

	rj := createResponseMsg(err)
	writeBack(w, r, rj)
}

// DELETEOrderHandler ...
func DELETEOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := parseReqBody(r, &order)

	if err == nil {
		err = order.Delete()
	}

	rj := createResponseMsg(err)

	writeBack(w, r, rj)
}
