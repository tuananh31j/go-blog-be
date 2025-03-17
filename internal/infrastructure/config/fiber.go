package config

import (
	"encoding/json"

	"nta-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig(apv string) fiber.Config {
	return fiber.Config{
		DisableStartupMessage: apv == "production",
		ServerHeader:          "Blog",
		AppName:               "Blog-API",
		ErrorHandler:          middleware.ErrorHandler,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	}
}
