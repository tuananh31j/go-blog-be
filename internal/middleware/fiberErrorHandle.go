package middleware

import (
	"nta-blog/internal/common"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(common.ErrInternal(err))
}
