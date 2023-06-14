package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	SECRET_JWT = ""
)

type AppConfig struct {
	DB_USERNAME           string
	DB_PASSWORD           string
	DB_HOSTNAME           string
	DB_PORT               int
	DB_NAME               string
	AWS_ACCESS_KEY_ID     string
	AWS_SECRET_ACCESS_KEY string
	jwtKey                string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.jwtKey = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DB_USERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DB_PASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DB_HOSTNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DB_PORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DB_NAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("AWS_ACCESS_KEY_ID"); found {
		app.AWS_ACCESS_KEY_ID = val
		isRead = false
	}
	if val, found := os.LookupEnv("AWS_SECRET_ACCESS_KEY"); found {
		app.AWS_SECRET_ACCESS_KEY = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		app.jwtKey = viper.Get("JWT_KEY").(string)
		app.DB_USERNAME = viper.Get("DBUSER").(string)
		app.DB_PASSWORD = viper.Get("DBPASS").(string)
		app.DB_HOSTNAME = viper.Get("DBHOST").(string)
		app.DB_PORT, _ = strconv.Atoi(viper.Get("DBPORT").(string))
		app.DB_NAME = viper.Get("DBNAME").(string)
		app.AWS_ACCESS_KEY_ID = viper.Get("AWS_ACCESS_KEY_ID").(string)
		app.AWS_SECRET_ACCESS_KEY = viper.Get("AWS_SECRET_ACCESS_KEY").(string)
	}

	SECRET_JWT = app.jwtKey
	fmt.Println("check", app.jwtKey)
	return &app
}
