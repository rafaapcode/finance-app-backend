package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-chi/jwtauth"
)

type conf struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	Jwt_Secret    string
	Jwt_ExpiresIn int
	tokenAuth     *jwtauth.JWTAuth
}

type jwtSecrets struct {
	Jwtexpires int
	TokenAuth  *jwtauth.JWTAuth
}

func loadConfig() conf {
	var cfg conf
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.WebServerPort = os.Getenv("WEB_SERVER_PORT")
	cfg.Jwt_Secret = os.Getenv("JWT_SECRET")
	expiresIn, err := strconv.Atoi(os.Getenv("JWT_EXPIRESIN"))
	if err != nil {
		panic(err.Error())
	}
	cfg.Jwt_ExpiresIn = expiresIn
	cfg.tokenAuth = jwtauth.New("HS256", []byte(cfg.Jwt_Secret), nil)

	return cfg
}

func GetDSN() string {
	conf := loadConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBName, conf.DBPort)
}

func GetPort() string {
	conf := loadConfig()
	return conf.WebServerPort
}

func GetJwtSecrets() jwtSecrets {
	conf := loadConfig()
	return jwtSecrets{
		TokenAuth:  conf.tokenAuth,
		Jwtexpires: conf.Jwt_ExpiresIn,
	}
}
