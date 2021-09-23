package api

import (
	"net/http"
)

// Get routes for GET requests of a defined path, the handle function and accepts a query parameter called "q".
// For large datasets, is a good practice use query parameter called "page", "limit" and  "offset".
func (srv *Service) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	srv.Router.HandleFunc(path, f).Methods(http.MethodGet)
}

// Post routes for POST requests of a defined path, the handle function.
func (srv *Service) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	srv.Router.HandleFunc(path, f).Methods(http.MethodPost)
}

// Put routes for PUT requests of a defined path, the handle function.
func (srv *Service) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	srv.Router.HandleFunc(path, f).Methods(http.MethodPut)
}

// Delete routes for DELETE requests of a defined path, the handle function.
func (srv *Service) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	srv.Router.HandleFunc(path, f).Methods(http.MethodDelete)
}
