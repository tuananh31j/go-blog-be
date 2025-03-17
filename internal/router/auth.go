package routes

import (
	authHttp "nta-blog/internal/domain/delivery/http/auth"
	"nta-blog/internal/infrastructure/config"
	"nta-blog/internal/lib/appctx"
	"nta-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func authRoutes(router fiber.Router, actx appctx.AppContext) {
	router.Post("/login", authHttp.Login(actx))
	router.Get("/google_callback", authHttp.GoogleLogin(actx))
	router.Post("/refresh-token", authHttp.RefreshToken(actx))
	router.Post("/logout", middleware.Authentication(config.Env.SecretAccessKey), authHttp.Logout(actx))
}
