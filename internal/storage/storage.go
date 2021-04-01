package storage

import (
	"database/sql"
	"log"
	"sync"

	"github.com/KovacDR/go-mysql-api/internal/config"
)


type Storage struct {
	DB *sql.DB
}

var (
	storage *Storage
	configuration *config.Config
	once sync.Once
)

func initDB() {
	configuration, _ = config.SetConfig()
	db, err := getConnection(configuration)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	
	storage = &Storage{
		DB: db,
	}

}

func New() (*Storage, error) {
	once.Do(initDB)

	return storage, nil
}
