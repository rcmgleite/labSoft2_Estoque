package main

import "fmt"

//Order is the struct that defines the purchase order
type Order struct {
	productList []Product
	size        int
	approved    bool
}

//NewOrder is a "constructor" for Order
func NewOrder(size int) *Order {
	return &Order{productList: make([]Product, size), size: 0, approved: false}
}

func (o *Order) addItem(p Product) {
	o.productList[o.size] = p
	o.size++
}

func (o *Order) removeItem(id int64) {
	//TODO
}

func (o *Order) approve() {
	fmt.Println("Passou Aqui!")
	o.approved = true
}
