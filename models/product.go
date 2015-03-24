package models

const (
	//FOOD ...
	FOOD = 1 << iota
	//CLEANING ...
	CLEANING
	//ROOMITENS ...
	ROOMITENS // towels, bed sheets
)

//Product struct that defines a product
type Product struct {
	ID           int
	Name         string `sql:"size:255"`
	Type         int
	Description  string `sql:"size:255"`
	CurrQuantity int
	MinQuantity  int
}

//NeedRefill verify if product need refill
func (p *Product) NeedRefill() bool {
	if p.CurrQuantity < p.MinQuantity {
		return true
	}
	return false
}
