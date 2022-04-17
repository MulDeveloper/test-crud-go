package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	r := chi.NewRouter()

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil

}

func (serv *Server) Close() error {
	//Close pending
	return nil
}

func (serv *Server) Start() {
	log.Printf("Server started, running at http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
