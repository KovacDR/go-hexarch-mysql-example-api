package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	v1 "github.com/KovacDR/go-mysql-api/internal/server/v1"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)


type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	r := chi.NewRouter()
	port = fmt.Sprintf(":%s", port)

	r.Mount("/api/v1", v1.New())

	hanlder := cors.AllowAll().Handler(r)

	serv := &http.Server{
		Addr: port,
		Handler: hanlder,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}