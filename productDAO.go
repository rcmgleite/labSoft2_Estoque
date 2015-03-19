package main

//ProductDAO ...
type ProductDAO struct {
	database *Database
}

func newProductDAO() *ProductDAO {
	return &ProductDAO{database: getDbInstance()}
}

// Save product on db
func (dao *ProductDAO) Save(p *Product) {
	dao.database.db.Create(p)
}

//Update product on db
func (dao *ProductDAO) Update(newProduct *Product) {
	database := getDbInstance()
	database.db.Save(newProduct)
}

// Delete product from db
func (dao *ProductDAO) Delete(p *Product) {
	database := getDbInstance()
	database.db.Delete(p)
}

//Retreive product from db
func (dao *ProductDAO) Retreive(ids ...int) []Product {
	if len(ids) == 0 {
		var products []Product
		dao.database.db.Find(&products)
		return products
	}
	return nil
	//TODO implementar para quando ids são usados
}
