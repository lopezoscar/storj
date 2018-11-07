package satellite

import "net/http"

type Server struct {
	db Database
}

func NewServer(db Database) *Server {
	return &Server{db}
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	server.db.Users().CreateUser(r.GetForm("x"), r.GetForm("y"))
}
