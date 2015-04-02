package models

const (
	//FOOD ...
	FOOD = 1 << iota
	//CLEANING ...
	CLEANING
	//ROOMITENS ...
	ROOMITENS // towels, bed sheets
)

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

//Retreive ... it uses the object and a plain query to execute sql cmds
func (p *Product) Retreive() ([]Product, error) {
	var query string
	if p.QueryParams != nil {
		query = buildQuery(p.QueryParams)
	}

	orderBy := p.QueryParams["order_by"]

	var products []Product
	var err error

	//Remove queryParams - gambi enquanto nao consigo tirar BaseModel das buscas
	p.QueryParams = nil
	if orderBy != "" {
		err = db.Order(orderBy).Where(*p).Find(&products, query).Error
	} else {
		err = db.Where(*p).Find(&products, query).Error
	}

	return products, err
}

//NeedRefill verify if product need refill
func (p *Product) NeedRefill() bool {
	if p.CurrQuantity < p.MinQuantity {
		return true
	}
	return false
}
