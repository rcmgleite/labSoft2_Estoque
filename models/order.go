package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/rcmgleite/labSoft2_Estoque/database"
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
	db := database.GetDatabase()
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
	db := database.GetDatabase()
	return db.Create(order).Error
}

// Update ...
func (order *Order) Update() error {
	db := database.GetDatabase()

	err := db.Save(order).Error
	if err != nil {
		return err
	}

	return order.send(comprasIP, "/order")
}

// Delete ...
func (order *Order) Delete() error {
	db := database.GetDatabase()

	err := db.Delete(order).Error
	return err
}

//GetOpenOrder ...
func GetOpenOrder(order *Order) error {
	db := database.GetDatabase()
	err := db.Where("approved = ?", false).First(order).Error
	if err != nil {
		return err
	}
	products := []Product{}
	err = db.Model(order).Related(&products, "Products").Error
	order.Products = products

	return err
}

// OpenOrderHasProduct ...
func OpenOrderHasProduct(product Product) (bool, error) {
	db := database.GetDatabase()

	order := Order{}
	err := GetOpenOrder(&order)
	if err != nil {
		return false, err
	}

	err = db.Model(order).Association("Products").Find(&product).Error

	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// RemoveProductFromOpenOrder from the existing opened order
func RemoveProductFromOpenOrder(product Product) error {
	db := database.GetDatabase()
	order := Order{}
	err := GetOpenOrder(&order)
	if err != nil {
		return err
	}
	return db.Model(order).Association("Products").Delete([]Product{product}).Error
}

// AddProductToOpenOrder to the existing opened order or creates a new order if needed
func AddProductToOpenOrder(product Product) error {
	var order Order
	err := GetOpenOrder(&order)
	if err != nil {
		if err.Error() == "record not found" {
			return order.createOrderAndAddProduct(product)
		}
		return err
	}
	return order.addProduct(product)
}

func (order *Order) createOrderAndAddProduct(product Product) error {
	db := database.GetDatabase()

	err := db.Create(order).Error
	if err != nil {
		return err
	}

	err = db.Model(order).Association("Products").Append([]Product{product}).Error
	if err != nil {
		return err
	}

	return nil
}

func (order *Order) addProduct(product Product) error {
	db := database.GetDatabase()

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
