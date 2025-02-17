package controller

import (
	"backend/middleware"
	"net/http"
)

func (server *Server) setJSON(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middleware.SetMiddlewareJSON(next)).Methods(method, "OPTIONS")

}

func (server *Server) initializeRoutes() {
	server.Router.Use(middleware.CORS)

	server.setJSON("/register-futsal", server.RegisterFutsal, "POST")
	server.setJSON("/get-all-futsal", server.GetAllFutsals, "GET")
	server.setJSON("/update-futsal/{id}", server.UpdateFutsal, "PUT")
	server.setJSON("/update-futsal-fields/{id}", server.UpdateFutsalFields, "PUT")
	server.setJSON("/delete-futsal/{id}", server.DeleteFutsal, "DELETE")

}
