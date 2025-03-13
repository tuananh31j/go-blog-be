package config

import (
	"nta-blog/internal/middleware"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig(apv string) fiber.Config {
	return fiber.Config{
		DisableStartupMessage: apv == "production",
		ServerHeader:          "Blog",
		AppName:               "Blog-API",
		ErrorHandler:          middleware.ErrorHandler,
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
	}
}
