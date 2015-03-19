package main

import "github.com/jinzhu/gorm"

//GenericDAO ...
type GenericDAO struct {
	db *gorm.DB
}

func newGenericDAO() *GenericDAO {
	dbFactory := getDbFactoryInstance("./estoque.db")
	return &GenericDAO{db: dbFactory.getDataBase()}
}

// Save product on db
func (dao *GenericDAO) Save(entity interface{}) {
	dao.db.Create(entity)
}

//Update product on db
func (dao *GenericDAO) Update(newEntity interface{}) {
	dao.db.Save(newEntity)
}

// Delete product from db
func (dao *GenericDAO) Delete(entity interface{}) {
	dao.db.Delete(entity)
}

//Retreive product from db
func (dao *GenericDAO) Retreive(ids ...int) interface{} {
	if len(ids) == 0 {
		var products []Product
		dao.db.Find(&products)
		return products
	}
	return nil
	//TODO implementar para quando ids são usados
}
