package server_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/rcmgleite/labSoft2_Estoque/models"
	"github.com/rcmgleite/labSoft2_Estoque/requestHelper"
)

func getJSON(object interface{}) ([]byte, error) {
	if object != nil {
		return json.Marshal(object)
	}
	return nil, nil
}

func TestCase1(t *testing.T) {
	var p models.Product
	p.Name = "test_product"
	p.Description = "test_descr"
	p.CurrQuantity = 100
	p.MinQuantity = 200
	p.Type = 2
	//Tests
	// 1) Request to add product to db
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	bJSON, err := getJSON(p)
	if err != nil {
		fmt.Println(err)
	} else {
		requestHelper.MakeRequest("POST", "http://127.0.0.1:8080/product", bJSON, headers)
	}

	// 2) Query for the product saved

	// 3) Verify if order was created

	// 4) Simple Query 1

	// 5) Simple Query 2

	// 6) Simple Query 3

	// 7) Simple Query 4

	// 8) Complex Query 1

	// 9) Complex Query 2

	// 10) Complex Query 3

	// 11) Complex Query 4

	// 12) Update 1

	// 12) Update 2

	// 12) Update 3

	// 12) Update 4

	// 12) Update 5
}
