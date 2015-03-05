package main

// ProductStore = mock for db
type ProductStore struct {
	Products map[string]Product
}

// NewProductStore = "constructor" for productStore
func NewProductStore() *ProductStore {
	return &ProductStore{Products: make(map[string]Product)}
}

//AddProduct add a product to db
func (ps *ProductStore) AddProduct(p Product) {
	ps.Products[p.Name] = p
}

//UpdateProduct update a product on db
func (ps *ProductStore) UpdateProduct(Product) {
	//TODO
}

//GetProduct retreive a product from db
func (ps *ProductStore) GetProduct(id int64) {
	//TODO
}

//DeleteProduct delete a product from db
func (ps *ProductStore) DeleteProduct(Product) {
	//TODO
}
