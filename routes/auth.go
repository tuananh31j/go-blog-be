package routes

import (
	"nta-blog/libs/appctx"
	authtransportfiber "nta-blog/modules/auth/transport/fiberAuth"

	"github.com/gofiber/fiber/v2"
)

func authRouter(router fiber.Router, actx appctx.AppContext) {
	router.Post("/login", authtransportfiber.Login(actx))
}
