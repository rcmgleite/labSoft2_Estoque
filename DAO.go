package main

import "github.com/jinzhu/gorm"

//DAO "Super class" for all DAOs on project
type DAO struct {
	db *gorm.DB
}

func newDAO(dbType, dbPath string) *DAO {
	dbFactory := getDbFactoryInstance(dbType)
	return &DAO{db: dbFactory.getDataBase(dbPath)}
}

// Save item on db
func (dao *DAO) Save(entity interface{}) error {
	return dao.db.Create(entity).Error
}

//Update item on db
func (dao *DAO) Update(newEntity interface{}) error {
	return dao.db.Save(newEntity).Error
}

// Delete item from db
func (dao *DAO) Delete(entity interface{}) error {
	return dao.db.Delete(entity).Error
}
