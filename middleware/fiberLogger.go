package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerConfig() fiber.Handler {
	return logger.New(logger.Config{
		Output:     os.Stdout,
		Format:     "${time} ${method} ${status} ${path} in ${latency}\n",
		TimeFormat: "15:04:05.00",
		Next:       nil,
	})
}
