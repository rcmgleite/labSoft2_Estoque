package main

import (
	"fmt"
	"net/http"
)

//HTTP METHODS
const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type route struct {
	path    string
	method  string
	handler http.HandlerFunc
}

//Router struct
type Router struct {
	routes map[string]*route
	// middleware filters TODO
	// filters []http.HandlerFunc
}

//NewRouter = constructor for router
func NewRouter() *Router {
	return &Router{routes: make(map[string]*route)}
}

//AddRoute adds a new route using path method and handler
func (r *Router) AddRoute(path string, method string, handler http.HandlerFunc) {
	switch method {
	case GET:
		r.doAddRoute(GET, path, handler)
		break
	case PUT:
		r.doAddRoute(PUT, path, handler)
		break

	case DELETE:
		r.doAddRoute(DELETE, path, handler)
		break

	case POST:
		r.doAddRoute(POST, path, handler)
		break
	}
}

//doAddRoute will add the specific route using method and string
func (r *Router) doAddRoute(method string, path string, handler http.HandlerFunc) {
	if r.routes[path+method] != nil {
		fmt.Printf("route with path '%s' with method '%s' already added. The second one will be ignored", path, method)
		return
	}

	route := &route{}
	route.path = path
	route.method = method
	route.handler = handler

	r.routes[path+method] = route
}

//ServeHTTP Implements interface http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	//TODO - add middleware handlers
	route := r.routes[rq.URL.Path+rq.Method]
	if route != nil {
		route.handler(w, rq)
	} else {
		//TODO - create not-found html
		http.Redirect(w, rq, "/", http.StatusNotFound)
	}
}
