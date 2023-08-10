package main

import (
	"create2-payment-gateway/controllers"
	"database/sql"
	"net/http"
)

// Router is a struct that holds the router configuration
// And some injected dependencies
type Router struct {
	db *sql.DB
}

// ServeHTTP is the main entry point for all requests
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error
	path := req.URL.Path

	switch path {
	case "/api/create":
		if req.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 Not Found"))
			return
		}
		err = controllers.CreatePayment(w, req, r.db)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR: " + err.Error()))
		return
	}
	return
}

// NewRouter returns a new Router instance
func NewRouter(db *sql.DB) *Router {
	return &Router{
		db: db,
	}
}
