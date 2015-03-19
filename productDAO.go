package main

import "github.com/jinzhu/gorm"

//ProductDAO ...
type ProductDAO struct {
	db *gorm.DB
}

func newProductDAO() *ProductDAO {
	dbFactory := getDbFactoryInstance("./estoque.db")
	return &ProductDAO{db: dbFactory.getDataBase()}
}

// Save product on db
func (dao *ProductDAO) Save(p *Product) {
	dao.db.Create(p)
}

//Update product on db
func (dao *ProductDAO) Update(newProduct *Product) {
	dao.db.Save(newProduct)
}

// Delete product from db
func (dao *ProductDAO) Delete(p *Product) {
	dao.db.Delete(p)
}

//Retreive product from db
func (dao *ProductDAO) Retreive(ids ...int) []Product {
	if len(ids) == 0 {
		var products []Product
		dao.db.Find(&products)
		return products
	}
	return nil
	//TODO implementar para quando ids são usados
}
