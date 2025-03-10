package common

import (
	"fmt"

	"nta-blog/libs/logger"

	jwt "github.com/dgrijalva/jwt-go"
)

func VerifyJWT(tokenString, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username := claims["id"].(string)
		return username, nil
	}

	return "", fmt.Errorf("failed to get claims")
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
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.ZeroLog.Debug().Msg(fmt.Sprintf("JWT error>>>%v", payload))
	}
	return t
}
