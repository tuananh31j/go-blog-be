package config

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	AppENV           string
	AppHost          string
	AppPort          string
	MongoURI         string
	RedisPort        string
	RedisHost        string
	RedisPass        string
	SecretAccessKey  string
	SecretRefreshKey string
}

var Env env

func init() {
	loadConfig()
	Env.AppHost = os.Getenv("APP_HOST")
	Env.AppPort = os.Getenv("APP_PORT")
	Env.MongoURI = os.Getenv("MONGO_URI")
	Env.RedisPort = os.Getenv("REDIS_PORT")
	Env.RedisHost = os.Getenv("REDIS_HOST")
	Env.RedisPass = os.Getenv("REDIS_PASS")
	Env.SecretAccessKey = os.Getenv("JWT_ACCESS_TOKEN_KEY")
	Env.SecretRefreshKey = os.Getenv("JWT_REFRESH_TOKEN_KEY")
	Env.AppENV = os.Getenv("APP_ENV")
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
