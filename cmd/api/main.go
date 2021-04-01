package main

import (
	"log"
	"os"
	"os/signal"

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
	_, err = storage.InitDB(cfg)
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
	
	go serv.Start()
	defer func() {
		if err = serv.Close(); err != nil {
			log.Fatal(err.Error())
			return
		}
	}()

	// wait 'til an interruption
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<- c
}