package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/rcmgleite/labSoft2_Estoque/models"
	"github.com/rcmgleite/labSoft2_Estoque/requestHelper"
)

//responseJSON
type responseJSON struct {
	ResponseBody interface{}
	Error        string
}

func getJSON(object interface{}) ([]byte, error) {
	if object != nil {
		return json.Marshal(object)
	}
	return nil, nil
}

func parseJSON(encodedJSON io.ReadCloser, object interface{}) {
	decoder := json.NewDecoder(encodedJSON)
	decoder.Decode(object)
}

// 1) Request to add product to db
func TestCase1(t *testing.T) {

	var p models.Product
	p.Name = "test_product"
	p.Description = "test_descr"
	p.CurrQuantity = 100
	p.MinQuantity = 200
	p.Type = 2

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	bJSON, err := getJSON(p)
	if err != nil {
		t.Error(err)
	} else {
		resp, err := requestHelper.MakeRequest("POST", "http://127.0.0.1:8080/product", bJSON, headers)
		if err != nil {
			t.Error(err)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		if !strings.Contains(string(body), "Product successfully saved") {
			t.Error(string(body))
		}
	}
}

// 2) Query for the product saved
func TestCase2(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	resp, err := requestHelper.MakeRequest("GET", "http://127.0.0.1:8080/product", nil, headers)
	if err != nil {
		t.Error(err)
	}
	var rj responseJSON
	parseJSON(resp.Body, &rj)
	bJSON, err := getJSON(rj.ResponseBody)
	var products []models.Product
	reader := bytes.NewReader(bJSON)
	decoder := json.NewDecoder(reader)
	decoder.Decode(&products)

	if len(products) != 1 {
		t.Error("Wrong number of products inserted on database")
	}

	if products[0].Name != "test_product" {
		t.Error("Wrong product.Name")
	}

}

// 3) Verify if order was created
func TestCase3(t *testing.T) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	resp, err := requestHelper.MakeRequest("GET", "http://127.0.0.1:8080/order", nil, headers)
	if err != nil {
		t.Error(err)
	}

	var rj responseJSON
	parseJSON(resp.Body, &rj)
	bJSON, err := getJSON(rj.ResponseBody)
	var order models.Order
	reader := bytes.NewReader(bJSON)
	decoder := json.NewDecoder(reader)
	decoder.Decode(&order)

	if len(order.Products) != 1 {
		t.Error("Wrong number of products on order")
	}

	if order.Products[0].Name != "test_product" {
		t.Error("Wrong product.Name")
	}

}

// 4) Simple Query
func TestCase4(t *testing.T) {

	// I) Creating more entries
	var p models.Product
	p.Name = "test_product2"
	p.Description = "test_descr2"
	p.CurrQuantity = 100
	p.MinQuantity = 200
	p.Type = 3

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	bJSON, _ := getJSON(p)
	requestHelper.MakeRequest("POST", "http://127.0.0.1:8080/product", bJSON, headers)

	p.Name = "test_product3"
	p.Description = "test_descr3"
	p.CurrQuantity = 300
	p.MinQuantity = 200
	p.Type = 4

	bJSON, _ = getJSON(p)
	requestHelper.MakeRequest("POST", "http://127.0.0.1:8080/product", bJSON, headers)

	p.Name = "test_product4"
	p.Description = "test_descr4"
	p.CurrQuantity = 100
	p.MinQuantity = 200
	p.Type = 5

	bJSON, _ = getJSON(p)
	requestHelper.MakeRequest("POST", "http://127.0.0.1:8080/product", bJSON, headers)

	// Query for this new objects added

	resp, err := requestHelper.MakeRequest("GET", "http://127.0.0.1:8080/product", nil, headers)
	if err != nil {
		t.Error(err)
	}
	var rj responseJSON
	parseJSON(resp.Body, &rj)
	bJSON, err = getJSON(rj.ResponseBody)
	var products []models.Product
	reader := bytes.NewReader(bJSON)
	decoder := json.NewDecoder(reader)
	decoder.Decode(&products)

	if len(products) != 4 {
		t.Error("Wrong number of products inserted on database")
	}

	if products[0].Name != "test_product" {
		t.Error("Wrong product.Name")
	}

	if products[1].Name != "test_product2" {
		t.Error("Wrong product.Name")
	}

	if products[2].Name != "test_product3" {
		t.Error("Wrong product.Name")
	}

	if products[3].Name != "test_product4" {
		t.Error("Wrong product.Name")
	}
}

// 5) Update
func TestCase5(t *testing.T) {
	// I) Update the item with id = 1
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	var toUpdate models.Product

	toUpdate.ID = 1
	toUpdate.Name = "test_product_updated"      //Original was test_product
	toUpdate.Description = "test_descr_updated" //Original was test_descr
	toUpdate.Type = 3                           //Original was 2
	toUpdate.CurrQuantity = 150                 //Original was 100
	toUpdate.MinQuantity = 250                  //Original was 200

	bJSON, err := getJSON(toUpdate)
	if err == nil {
		_, err := requestHelper.MakeRequest("PUT", "http://127.0.0.1:8080/product", bJSON, headers)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Query the item with id = 1 and verify if the update worked
	resp, err := requestHelper.MakeRequest("GET", "http://127.0.0.1:8080/product?ID=1", nil, headers)
	if err != nil {
		t.Error(err)
	}

	var rj responseJSON
	parseJSON(resp.Body, &rj)
	bJSON, err = getJSON(rj.ResponseBody)
	var products []models.Product
	reader := bytes.NewReader(bJSON)
	decoder := json.NewDecoder(reader)
	decoder.Decode(&products)

	if len(products) != 1 {
		t.Error("Wrong length for products")
	}

	p := products[0]
	if p.Name != "test_product_updated" || p.Description != "test_descr_updated" || p.Type != 3 || p.CurrQuantity != 150 || p.MinQuantity != 250 {
		t.Error("Update didn't worked")
	}
}
