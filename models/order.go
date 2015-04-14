package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/rcmgleite/labSoft2_Estoque/requestHelper"
)

// OrderToSend ...
type OrderToSend struct {
	Products []ProductToSend `json:"produtos"`
}

//Order is the struct that defines the purchase order
type Order struct {
	BaseModel `sql:"-" json:",omitempty"` // Ignore this field
	ID        int
	Products  []Product `gorm:"many2many:order_products;"`
	Valor     int       `json:"valor" sql:"-"`
	Approved  bool
}

// GetByID ...
func (order *Order) GetByID(id int) error {
	err := db.Where("id = ?", id).First(order).Error
	if err != nil {
		return err
	}
	products := []Product{}
	err = db.Model(order).Related(&products, "Products").Error
	order.Products = products

	return err
}

//Save ..
func (order *Order) Save() error {
	return db.Create(order).Error
}

// Update ...
func (order *Order) Update() error {
	err := db.Save(order).Error
	if err != nil {
		return err
	}
	return order.send(comprasIP, "/order")
}

// Delete ...
func (order *Order) Delete() error {
	return db.Delete(order).Error
}

//GetOpenOrder ...
func (order *Order) GetOpenOrder() error {
	err := db.Where("approved = ?", false).First(order).Error
	if err != nil {
		return err
	}
	products := []Product{}
	err = db.Model(order).Related(&products, "Products").Error
	order.Products = products

	return err
}

// AddProduct ...
func (order *Order) AddProduct(product Product) error {
	err := order.GetOpenOrder()
	if err != nil {
		if err.Error() == "record not found" {
			return order.createOrderAndAddProduct(product)
		}
		return err
	}
	return order.addProduct(product)
}

func (order *Order) createOrderAndAddProduct(product Product) error {
	//Creates a single transaction to create and add new product to order
	tx := db.Begin()

	err := tx.Create(order).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(order).Association("Products").Append([]Product{product}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (order *Order) addProduct(product Product) error {
	return db.Model(order).Association("Products").Append([]Product{product}).Error
}

// FIXME - MOVE THIS FUNCTION TO A PROPER HELPER
func getJSON(object interface{}) ([]byte, error) {
	if object != nil {
		return json.Marshal(object)
	}
	return nil, nil
}

//send = Function that sends the new order to compras module
func (order *Order) send(dstIP string, dstPath string) error {
	err := order.GetByID(order.ID)

	if err == nil {
		headers := make(map[string]string)
		headers["Content-Type"] = "application/json"

		pToSend := make([]ProductToSend, len(order.Products))
		for index, value := range order.Products {
			pToSend[index].ProductID = value.ID
			pToSend[index].Quantidade = value.MinQuantity - value.CurrQuantity
			pToSend[index].Valor = 10
		}

		orderToSend := &OrderToSend{Products: pToSend}
		bJSON, err := getJSON(orderToSend)
		if err == nil {
			resp, err := requestHelper.MakeRequest("POST", dstIP+dstPath, bJSON, headers)
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
