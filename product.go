package main

//Product struct that defines a product
type Product struct {
	name         string
	description  string
	currQuantity int64
	minQuantity  int64
}

func (p *Product) needRefill() bool {
	if p.currQuantity < p.minQuantity {
		return true
	}
	return false
}
