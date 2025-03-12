package routes

import (
	authHttp "nta-blog/internal/domain/delivery/http/auth"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func authRouter(router fiber.Router, actx appctx.AppContext) {
	router.Post("/login", authHttp.Login(actx))
	router.Get("/google_callback", authHttp.GoogleLogin(actx))
}
