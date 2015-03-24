package client

import (
	"fmt"
	"net/http"

	"github.com/rcmgleite/labEngSoft_Estoque/router"
)

//Run = main for client
func Run() {
	fmt.Println("Running client")
	r := router.NewRouter()
	r.AddRoute("/client", router.GET, defaultHandler)
	http.Handle("/client", r)

	http.ListenAndServe(":8081", nil)
}
