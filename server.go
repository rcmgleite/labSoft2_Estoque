package main

import (
	"fmt"
	"net/http"

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
	r.AddRoute("/order", router.PUT, PUTOrderHandler)
	r.AddRoute("/order", router.DELETE, DELETEOrderHandler)

	http.Handle("/", r)

	fmt.Println("Server running on port: 8080")

	http.ListenAndServe(":8080", nil)
}
