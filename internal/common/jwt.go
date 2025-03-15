package common

import (
	"fmt"

	"nta-blog/internal/lib/logger"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

type Payload struct {
	Id   string `json:"id"`
	Role string `json:"role"`
}

func VerifyJWT(tokenString, secret string) (*Payload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	var payload Payload
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		mapstructure.Decode(claims, &payload)

		return &payload, nil
	}

	return nil, fmt.Errorf("failed to get claims")
}

func GenerateJWT(secret string, payload map[string]string, exp int64) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logger.ZeroLog.Debug().Msg("JWT something wrong!")
	}

	for key, value := range payload {
		claims[key] = value
	}
	claims["exp"] = exp
	claims["batchedova"] = "Em! Tính làm gì vậy em?"
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.ZeroLog.Debug().Msg(fmt.Sprintf("JWT error>>>%v", payload))
	}
	return t
}
