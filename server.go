package main

import (
	"net/http"

	"github.com/rcmgleite/labEngSoft_Estoque/client"
	"github.com/rcmgleite/labEngSoft_Estoque/router"
)

func main() {
	r := router.NewRouter()

	// /product
	r.AddRoute("/product", router.GET, GETProductHandler)
	r.AddRoute("/product", router.POST, POSTProductHandler)
	r.AddRoute("/product", router.PUT, PUTProductHandler)
	r.AddRoute("/product", router.DELETE, DELETEProductHandler)

	// /order
	r.AddRoute("/order", router.GET, GETOrderHandler)
	r.AddRoute("/order", router.POST, POSTOrderHandler)
	r.AddRoute("/order", router.PUT, PUTOrderHandler)
	r.AddRoute("/order", router.DELETE, DELETEOrderHandler)

	http.Handle("/", r)

	client.Run()
	http.ListenAndServe(":8080", nil)
}
