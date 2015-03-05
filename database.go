package main

type Database struct {
	Produtos map[string]Produto
}

func NewDatabase() *Database {
	return &Database{Produtos: make(map[string]Produto)}
}

func (ps *Database) AddProduto(p Produto) {
	ps.Produtos[p.Name] = p
}
