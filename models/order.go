package models

// var _db = database.GetDbFactoryInstance("sqlite3").GetDatabase("./estoque.db")

//Order is the struct that defines the purchase order
type Order struct {
	BaseModel `sql:"-"` // Ignore this field
	ID        int
	Products  []Product `gorm:"many2many:order_products;"`
	Approved  bool
}

//Save ..
func (order *Order) Save() error {
	return db.Create(order).Error
}

// Update ...
func (order *Order) Update() error {
	return db.Save(order).Error
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
	return db.Model(order).Association("Products").Append([]Product{product}).Error
}
