package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

func parseRequestFormProduct(r *http.Request, p *Product) bool {
	val := reflect.ValueOf(p).Elem()

	//Starts from 1 to ignore ID
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		if field.Name == "ID" {
			continue
		}
		r.FormValue(field.Name)
		switch reflect.ValueOf(p).Elem().Field(i).Kind() {
		case reflect.Int:
			value, err := strconv.ParseInt(r.FormValue(field.Name), 0, 64)
			if err != nil {
				fmt.Println(err)
				fmt.Println(field.Name)
				return false
			}
			reflect.ValueOf(p).Elem().Field(i).SetInt(value)

			break
		case reflect.String:
			reflect.ValueOf(p).Elem().Field(i).SetString(r.FormValue(field.Name))
		}
	}
	return true
}
