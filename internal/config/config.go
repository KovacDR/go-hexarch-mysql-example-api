package config

import (
	"github.com/spf13/viper"
)


type Config struct {
	PORT        string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_SSLMODE  string
	DB_HOST     string
	DB_PORT 	string
}

func SetConfig() (*Config, error) {
	var configuration Config

	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigType("yaml")
	cfg.AddConfigPath(".")

	err := cfg.ReadInConfig()
	if err != nil {
		return &configuration, err
	}
	
	
	err = cfg.Unmarshal(&configuration)
	if err != nil {
		return &configuration, err
	}

	return &configuration, nil
}