package main

//Produto struct that defines a Produto
type Produto struct {
	Name         string
	CurrQuantity int64
	MinQuantity  int64
}

func (p *Produto) precisaRepor() bool {
	if p.CurrQuantity < p.MinQuantity {
		return true
	}
	return false
}
