package main

import "github.com/rcmgleite/labEngSoft_Estoque/models"

//OrderDAO ... subclass of DAO
type OrderDAO struct {
	*DAO
}

func newOrderDAO() *OrderDAO {
	dao := newDAO("sqlite3", "./estoque.db")
	return &OrderDAO{dao}
}

//Retreive ...
func (dao *OrderDAO) Retreive(toFill *models.Order) error {
	products := []models.Product{}
	err := dao.db.Debug().Model(toFill).Related(&products, "Products").Error
	toFill.Products = products

	return err
}
