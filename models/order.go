package models

//Order is the struct that defines the purchase order
type Order struct {
	ID       int
	Products []Product `gorm:"many2many:order_products;"`
	Approved bool
}

//NewOrder is a "constructor" for Order
func NewOrder() *Order {
	return &Order{}
}

func (o *Order) addItem(p Product) {
	o.Products = append(o.Products, p)
}

func (o *Order) removeItem(id int64) {
	//TODO
}

func (o *Order) approve() {
	o.Approved = true
}
