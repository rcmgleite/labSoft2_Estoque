package models

import "github.com/rcmgleite/labSoft2_Estoque/database"

var db = database.GetDbFactoryInstance("sqlite3").GetDatabase("./estoque.db")

//BaseModel struct for all models
type BaseModel struct {
	queryParams map[string]string
}

//Identifiers for query
const (
	LAST_MODIFIED_GT = "last_modified_gt"
	CREATED_AT_GT    = "created_at_gt"
)

//Aux functions
func buildQuery(queryMap map[string]string) string {
	return ""
}
