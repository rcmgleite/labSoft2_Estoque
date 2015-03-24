package main

import (
	"github.com/jinzhu/gorm"
	"github.com/rcmgleite/labEngSoft_Estoque/models"
)

//GenericDAO ...
type GenericDAO struct {
	db *gorm.DB
}

func newGenericDAO() *GenericDAO {
	dbFactory := getDbFactoryInstance("sqlite3")
	return &GenericDAO{db: dbFactory.getDataBase("./estoque.db")}
}

// Save product on db
func (dao *GenericDAO) Save(entity interface{}) error {
	return dao.db.Create(entity).Error
}

//Update product on db
func (dao *GenericDAO) Update(newEntity interface{}) error {
	return dao.db.Save(newEntity).Error
}

// Delete product from db
func (dao *GenericDAO) Delete(entity interface{}) error {
	return dao.db.Delete(entity).Error
}

//Retreive product from db
func (dao *GenericDAO) Retreive(ids ...int) interface{} {
	if len(ids) == 0 {
		var products []models.Product
		dao.db.Find(&products)
		return products
	}
	return nil
	//TODO implementar para quando ids são usados
}
