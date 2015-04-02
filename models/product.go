package models

const (
	//FOOD ...
	FOOD = 1 << iota
	//CLEANING ...
	CLEANING
	//ROOMITENS ...
	ROOMITENS // towels, bed sheets
)

// var db = database.GetDbFactoryInstance("sqlite3").GetDatabase("./estoque.db")

//Product struct that defines a product
type Product struct {
	BaseModel    `sql:"-"` // Ignore this field
	ID           int
	Name         string `sql:"size:255"`
	Type         int
	Description  string `sql:"size:255"`
	CurrQuantity int
	MinQuantity  int
}

//Save ..
func (p *Product) Save() error {
	return db.Create(p).Error
}

// Update ...
func (p *Product) Update() error {
	return db.Save(p).Error
}

// Delete ...
func (p *Product) Delete() error {
	return db.Delete(p).Error
}

//Retreive ...
func (p *Product) Retreive() ([]Product, error) {

	var query string
	if p.queryParams != nil {
		query = buildQuery(p.queryParams)
	}

	var products []Product
	err := db.Where(*p).Find(&products, query).Error

	return products, err
}

//NeedRefill verify if product need refill
func (p *Product) NeedRefill() bool {
	if p.CurrQuantity < p.MinQuantity {
		return true
	}
	return false
}
