package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)


type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	r := chi.NewRouter()
	hanlder := cors.AllowAll().Handler(r)
	port = fmt.Sprintf(":%s", port)

	serv := &http.Server{
		Addr: port,
		Handler: hanlder,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

func (serv *Server) Close() error {
	// close all resources
	return nil
}

func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}