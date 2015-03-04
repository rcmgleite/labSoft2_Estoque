package main

import "net/http"

//Router struct
type Router struct {
}

//Route method.
func (router *Router) Route() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/product", productHandler)
}
