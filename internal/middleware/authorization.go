package middleware

import (
	"strconv"
	"strings"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func Authorization(apctx appctx.AppContext, secret string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		payload, err := common.VerifyJWT(tokenString, secret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewUnauthorized(err, "Unauthorized", "Token is invalid"))
		}

		role, err := strconv.Atoi(payload.Role)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewUnauthorized(err, "Unauthorized", "Role is invalid"))
		}
		isAdmin := cnst.TRoleAccount(role) == cnst.Role.Admin
		if !isAdmin {
			return c.Status(fiber.StatusUnauthorized).JSON(common.NewUnauthorized(err, "Unauthorized", "Not an admin"))
		}
		return c.Next()
	}
}
