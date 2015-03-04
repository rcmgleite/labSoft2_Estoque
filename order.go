package main

//Order is the struct that defines the purchase order
type Order struct {
	productList map[string]Product
	size        int
	approved    bool
}

//NewOrder is a "constructor" for Order
func NewOrder(size int) *Order {
	return &Order{productList: make(map[string]Product, size), size: 0, approved: false}
}

func (o *Order) addItem(p Product) {
	o.productList[p.name] = p
	o.size++
}

func (o *Order) removeItem(id int64) {
	//TODO
}

func (o *Order) approve() {
	o.approved = true
}
