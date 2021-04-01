package storage

import (
	"database/sql"

	"github.com/KovacDR/go-mysql-api/internal/config"
)


type Storage struct {
	DB *sql.DB
}

var (
	storage *Storage
)

func InitDB(config *config.Config) (*Storage, error) {
	db, err := getConnection(config)
	if err != nil {
		return &Storage{}, err
	}
	
	storage = &Storage{
		DB: db,
	}

	return storage, nil
}
