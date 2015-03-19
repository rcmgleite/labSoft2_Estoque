package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

// BuildStructFromForm will iterate over request params and build the corresponding struct passed as parameter
// The struct fields must have the same names as the form
func BuildStructFromForm(r *http.Request, genericStruct interface{}) bool {
	val := reflect.ValueOf(genericStruct).Elem()

	//Starts from 1 to ignore ID
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		if field.Name == "ID" {
			continue
		}
		r.FormValue(field.Name)
		switch reflect.ValueOf(genericStruct).Elem().Field(i).Kind() {
		case reflect.Int:
			value, err := strconv.ParseInt(r.FormValue(field.Name), 0, 64)
			if err != nil {
				fmt.Println(err)
				fmt.Println(field.Name)
				return false
			}
			reflect.ValueOf(genericStruct).Elem().Field(i).SetInt(value)

			break
		case reflect.String:
			reflect.ValueOf(genericStruct).Elem().Field(i).SetString(r.FormValue(field.Name))
		}
	}
	return true
}
