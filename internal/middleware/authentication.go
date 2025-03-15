package middleware

import (
	"errors"
	"strings"

	"nta-blog/internal/common"
	"nta-blog/internal/lib/logger"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Authentication(accessSecret string) func(C *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if authHeader == "" || tokenString == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewUnauthorized(errors.New("Unauthorized"), "Unauthorized", "Token is required"))
		} else {
			payload, err := common.VerifyJWT(tokenString, accessSecret)
			if err != nil {
				logger.ZeroLog.Debug().Err(err).Msg("Cannot verify token!")
				return c.Status(fiber.StatusUnauthorized).JSON(common.NewUnauthorized(err, "Unauthorized", "Token is invalid"))
			}

			objId, err := primitive.ObjectIDFromHex(payload.Id)
			if err != nil {
				logger.ZeroLog.Debug().Err(err).Msg("Cannot convert userId to ObjectID")
				return c.Status(fiber.StatusUnauthorized).JSON(common.NewUnauthorized(err, "Unauthorized", "Token is invalid"))
			}

			c.Locals("userId", objId)
			return c.Next()
		}
	}
}
