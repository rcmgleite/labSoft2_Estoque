package main

//ProductDAO ... subclass of DAO
type ProductDAO struct {
	*DAO
}

func newProductDAO() *ProductDAO {
	dao := newDAO("sqlite3", "./estoque.db")
	return &ProductDAO{dao}
}

//Retreive item from db
func (dao *ProductDAO) Retreive(toFill interface{}, ids ...int) error {
	if len(ids) == 0 {
		return dao.db.Find(toFill).Error
	}
	return nil
	//TODO implementar para quando ids são usados
}
