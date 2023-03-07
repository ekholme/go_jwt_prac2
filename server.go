package jwt_prac2

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router    *mux.Router
	Srvr      *http.Server
	UserStore *UserStore
}

func NewServer(r *mux.Router, us *UserStore) *Server {
	listenAddr := ":8080"

	return &Server{
		Router: r,
		Srvr: &http.Server{
			Addr: listenAddr,
		},
		UserStore: us,
	}
}
