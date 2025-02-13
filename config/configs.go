package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-chi/jwtauth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type conf struct {
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	WebServerPort    string
	Jwt_Secret       string
	Jwt_ExpiresIn    int
	GOOGLE_CLIENT_ID string
	GOOGLE_SECRET    string
	tokenAuth        *jwtauth.JWTAuth
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
	cfg.GOOGLE_CLIENT_ID = os.Getenv("GOOGLE_CLIENT_ID")
	cfg.GOOGLE_SECRET = os.Getenv("GOOGLE_SECRET")

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

func OauthConfig() *oauth2.Config {
	conf := loadConfig()
	return &oauth2.Config{
		ClientID:     conf.GOOGLE_CLIENT_ID,
		ClientSecret: conf.GOOGLE_SECRET,
		RedirectURL:  "http://localhost:3003/users/auth/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}
