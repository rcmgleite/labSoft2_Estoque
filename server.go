package main

import "net/http"

func main() {
	r := NewRouter()
	r.AddRoute("/", GET, defaultHandler)

	// /product
	r.AddRoute("/product", GET, GETProductHandler)
	r.AddRoute("/product", POST, POSTProductHandler)
	r.AddRoute("/product", PUT, PUTProductHandler)
	r.AddRoute("/product", DELETE, DELETEProductHandler)

	// /order
	r.AddRoute("/order", GET, GETOrderHandler)
	r.AddRoute("/order", POST, POSTOrderHandler)
	r.AddRoute("/order", PUT, PUTOrderHandler)
	r.AddRoute("/order", DELETE, DELETEOrderHandler)

	// TODO - fix requests to css and js files
	// r.AddRoute("/css", GET, GETcssHandler)
	// r.AddRoute("/js", GET, GETjsHandler)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
