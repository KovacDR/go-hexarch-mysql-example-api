package main

import (
	"log"

	"github.com/KovacDR/go-mysql-api/internal/config"
	"github.com/KovacDR/go-mysql-api/internal/server"
	"github.com/KovacDR/go-mysql-api/internal/storage"
)


func main() {
	// load configuration
	cfg, err := config.SetConfig()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	
	// init database
	storage.New()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// star the server
	serv, err := server.New(cfg.PORT)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	
	serv.Start()
}