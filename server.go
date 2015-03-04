package main

import "net/http"

func main() {
	router := Router{}
	router.Route()
	http.ListenAndServe(":8080", nil)
}
