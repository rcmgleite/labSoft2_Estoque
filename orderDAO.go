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

//GetOpenOrder ...
func (dao *OrderDAO) GetOpenOrder(toFill *models.Order) error {
	err := dao.db.Where("approved = ?", false).First(toFill).Error
	if err != nil {
		return err
	}
	products := []models.Product{}
	err = dao.db.Model(toFill).Related(&products, "Products").Error
	toFill.Products = products

	return err
}

// AddProduct ...
func (dao *OrderDAO) AddProduct(product models.Product) error {
	var toUpdate models.Order
	err := dao.GetOpenOrder(&toUpdate)
	if err == nil {
		return dao.db.Model(&toUpdate).Association("Products").Append([]models.Product{product}).Error
	}
	if err.Error() == "record not found" {
		err = dao.Save(&toUpdate)
		return dao.AddProduct(product)
	}
	return err
}
