package config

import (
	"fmt"
	"os"
)

type conf struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
}

func loadConfig() conf {
	var cfg conf
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.WebServerPort = os.Getenv("WEB_SERVER_PORT")

	return cfg
}

func GetDSN() string {
	conf := loadConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBName, conf.DBPort)
}

func GetPort() string {
	conf := loadConfig()
	return conf.WebServerPort
}
