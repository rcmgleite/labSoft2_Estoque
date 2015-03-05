package main

//Order is the struct that defines the purchase order
type Order struct {
	ProductList map[string]Product
	Approved    bool
}

//NewOrder is a "constructor" for Order
func NewOrder() *Order {
	return &Order{ProductList: make(map[string]Product), Approved: false}
}

func (o *Order) addItem(p Product) {
	o.ProductList[p.Name] = p
}

func (o *Order) removeItem(id int64) {
	//TODO
}

func (o *Order) approve() {
	o.Approved = true
}
