package server

import (
	"log"
	"net/http"
	"time"

	v0 "github.com/MulDeveloper/go-test-crud/internal/server/v0"
	"github.com/go-chi/chi"
)

type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	r := chi.NewRouter()

	// API routes v0.
	r.Mount("/api/v0", v0.New())

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
