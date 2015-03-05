package main

type Pedido struct {
	ProdutoList map[string]Produto
	Approved    bool
}

func NewPedido() *Pedido {
	return &Pedido{ProdutoList: make(map[string]Produto), Approved: false}
}

func (o *Pedido) addItem(p Produto) {
	o.ProdutoList[p.Name] = p
}

func (o *Pedido) approve() {
	o.Approved = true
}
