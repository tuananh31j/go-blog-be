package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	AppENV   string
	AppHost  string
	AppPort  string
	MongoURI string

	RedisURL string

	SecretAccessKey  string
	SecretRefreshKey string

	GoogleClientId       string
	GoogleClientSecret   string
	GoogleRedirectServer string

	CloudinaryAPIKey    string
	CloudinaryAPISecret string
	CloudinaryName      string

	NextJSRedirectOauth string
	AllowOrigin         string
}

var Env env

func init() {
	loadConfig()
	Env.AppHost = os.Getenv("APP_HOST")
	Env.AppPort = os.Getenv("APP_PORT")

	Env.MongoURI = os.Getenv("MONGO_URI")

	Env.RedisURL = os.Getenv("REDIS_URL")

	Env.SecretAccessKey = os.Getenv("JWT_ACCESS_TOKEN_KEY")
	Env.SecretRefreshKey = os.Getenv("JWT_REFRESH_TOKEN_KEY")

	Env.AppENV = os.Getenv("APP_ENV")

	Env.GoogleClientId = os.Getenv("GOOGLE_CLIENT_ID")
	Env.GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	Env.GoogleRedirectServer = os.Getenv("GOOGLE_REDIRECT_SERVER")

	Env.CloudinaryAPIKey = os.Getenv("CLOUDINARY_API_KEY")
	Env.CloudinaryAPISecret = os.Getenv("CLOUDINARY_API_SECRET")
	Env.CloudinaryName = os.Getenv("CLOUDINARY_CLOUD_NAME")

	Env.NextJSRedirectOauth = os.Getenv("NEXTJS_REDIRECT_OAUTH")
	Env.AllowOrigin = os.Getenv("ALLOWS_ORIGIN")

	GoogleConfig(Env.GoogleClientId, Env.GoogleClientSecret, Env.GoogleRedirectServer)
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
