package main

import "net/http"

func main() {
	http.HandleFunc("/produto", produtoHandler)
	http.HandleFunc("/pedido", pedidoHandler)
	http.ListenAndServe(":8080", nil)
}
