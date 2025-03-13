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

	GoogleClientId     string
	GoogleClientSecret string

	CloudinaryAPIKey    string
	CloudinaryAPISecret string
	CloudinaryName      string
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

	Env.GoogleClientId = os.Getenv("GOOGLE_CLIENT_ID")
	Env.GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

	Env.CloudinaryAPIKey = os.Getenv("CLOUDINARY_API_KEY")
	Env.CloudinaryAPISecret = os.Getenv("CLOUDINARY_API_SECRET")
	Env.CloudinaryName = os.Getenv("CLOUDINARY_CLOUD_NAME")

	GoogleConfig(Env.GoogleClientId, Env.GoogleClientSecret)
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
