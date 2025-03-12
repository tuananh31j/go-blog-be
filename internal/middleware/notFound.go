package middleware

import (
	"errors"

	"nta-blog/internal/common"

	"github.com/gofiber/fiber/v2"
)

func NotFound(fctx *fiber.Ctx) error {
	return fctx.Status(fiber.StatusNotFound).JSON(common.ErrNotFound(errors.New("The requested path does not exist")))
}
