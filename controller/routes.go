package controller

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/register-futsal", server.RegisterFutsal).Methods("POST")
}
