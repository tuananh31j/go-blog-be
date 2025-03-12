package routes

import (
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, actx appctx.AppContext) {
	routerV1 := app.Group("/api/v1")

	auth := routerV1.Group("/auth")
	authRouter(auth, actx)
}
