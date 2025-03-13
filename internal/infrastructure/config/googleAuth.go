package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleLoginConfig oauth2.Config
}

var GuConfig Config

func GoogleConfig(clientId, clientSecret string) {
	GuConfig.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  "http://localhost:5000/api/v1/auth/google_callback",
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
