package models

import (
	"github.com/rcmgleite/labEngSoft_Estoque/database"
)

var _db = database.GetDbFactoryInstance("sqlite3").GetDatabase("./estoque.db")

//Order is the struct that defines the purchase order
type Order struct {
	ID          int
	Products    []Product `gorm:"many2many:order_products;"`
	Approved    bool
	queryParams map[string]string `sql:"-"` // Ignore this field
}

//Save ..
func (order *Order) Save() error {
	return _db.Create(order).Error
}

// Update ...
func (order *Order) Update() error {
	return _db.Save(order).Error
}

// Delete ...
func (order *Order) Delete() error {
	return _db.Delete(order).Error
}

//GetOpenOrder ...
func (order *Order) GetOpenOrder() error {
	err := _db.Where("approved = ?", false).First(order).Error
	if err != nil {
		return err
	}
	products := []Product{}
	err = _db.Model(order).Related(&products, "Products").Error
	order.Products = products

	return err
}

// AddProduct ...
func (order *Order) AddProduct(product Product) error {
	return _db.Model(order).Association("Products").Append([]Product{product}).Error
	// err := order.GetOpenOrder()
	// if err == nil {
	// 	return
	// }
	// if err.Error() == "record not found" {
	// 	err = order.Save()
	// 	return order.AddProduct(product)
	// }
	// return err
}
