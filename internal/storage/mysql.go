package storage

import (
	"database/sql"
	"fmt"

	"github.com/KovacDR/go-mysql-api/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func getConnection(config *config.Config) (*sql.DB, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	return sql.Open("mysql", uri)
}