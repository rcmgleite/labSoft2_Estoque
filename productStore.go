package main

// ProductStore = mock for db
type ProductStore struct {
	Size     int
	Products []Product
}

// NewProductStore = "constructor" for productStore
func NewProductStore(size int) *ProductStore {
	return &ProductStore{Products: make([]Product, size), Size: 0}
}

//AddProduct add a product to db
func (ps *ProductStore) AddProduct(p *Product) {
	ps.Products[ps.Size] = *p
	ps.Size++
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
