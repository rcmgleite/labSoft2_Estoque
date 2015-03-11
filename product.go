package main

//Product struct that defines a product
type Product struct {
	Name         string
	Type         int
	Description  string
	CurrQuantity int64
	MinQuantity  int64
}

func (p *Product) needRefill() bool {
	if p.CurrQuantity < p.MinQuantity {
		return true
	}
	return false
}
