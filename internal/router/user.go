package routes

import (
	userHttp "nta-blog/internal/domain/delivery/http/user"
	"nta-blog/internal/infrastructure/config"
	"nta-blog/internal/lib/appctx"
	"nta-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func userRoutes(router fiber.Router, actx appctx.AppContext) {
	router.Use(middleware.Authentication(config.Env.SecretAccessKey))

	router.Get("/me", userHttp.GetMe(actx))
}
