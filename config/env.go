package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	AppHost   string
	AppPort   string
	MongoURI  string
	RedisPort string
	RedisHost string
	RedisPass string
)

func init() {
	loadConfig()
	AppHost = os.Getenv("APP_HOST")
	AppPort = os.Getenv("APP_PORT")
	MongoURI = os.Getenv("MONGO_URI")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPass = os.Getenv("REDIS_PASS")
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
