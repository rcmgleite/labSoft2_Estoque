package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/rcmgleite/labEngSoft_Estoque/models"
)

type responseJSON struct {
	Msg string
}

var dao = newGenericDAO()

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

// GETProductHandler ...
func GETProductHandler(w http.ResponseWriter, r *http.Request) {
	writeBack(w, r, dao.Retreive())
}

// POSTProductHandler ...
func POSTProductHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	decoder := json.NewDecoder(r.Body)
	var p models.Product
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println(err)
	} else {
		err = dao.Save(&p)
		if p.NeedRefill() {
			fmt.Println("will need refill")
		}
		fmt.Println(p)
	}

	var msg string
	if err != nil {
		msg = err.Error()
	} else {
		msg = "Sucess"
	}
	rj := responseJSON{Msg: msg}

	writeBack(w, r, rj)
}

// PUTProductHandler ...
func PUTProductHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	decoder := json.NewDecoder(r.Body)
	var p models.Product
	err := decoder.Decode(&p)

	if err == nil {
		err = dao.Update(&p)
	}

	var msg string
	if err != nil {
		msg = err.Error()
	} else {
		msg = "Sucess"
	}

	rj := responseJSON{Msg: msg}

	writeBack(w, r, rj)
}

// DELETEProductHandler ...
func DELETEProductHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	decoder := json.NewDecoder(r.Body)
	var p models.Product
	err := decoder.Decode(&p)

	if err == nil {
		err = dao.Delete(&p)
	}

	var msg string
	if err != nil {
		msg = err.Error()
	} else {
		msg = "Sucess"
	}

	rj := responseJSON{Msg: msg}

	writeBack(w, r, rj)
}

// GETOrderHandler ...
func GETOrderHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	// t, _ := template.ParseFiles("views/html/order.html")
	// t.Execute(w, nil)
}

// POSTOrderHandler ...
func POSTOrderHandler(w http.ResponseWriter, r *http.Request) {
	//TODO
	// uOrder.approve()
	// http.Redirect(w, r, "/order", http.StatusFound)
}

// PUTOrderHandler ...
func PUTOrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUTOrderHandler")
	//TODO
}

// DELETEOrderHandler ...
func DELETEOrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETEOrderHandler")
	//TODO
}
