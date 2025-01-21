package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type conf struct {
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
}

func loadConfig() (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}

func GetDSN() (string, error) {
	conf, err := loadConfig()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBName, conf.DBPort), nil
}

func GetPort() (string, error) {
	conf, err := loadConfig()
	if err != nil {
		panic(err)
	}

	return conf.WebServerPort, nil
}
