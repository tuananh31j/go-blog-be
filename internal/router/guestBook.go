package routes

import (
	guestBookHttp "nta-blog/internal/domain/delivery/http/guestBook"
	"nta-blog/internal/infrastructure/config"
	"nta-blog/internal/lib/appctx"
	"nta-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func guestBookRoutes(router fiber.Router, actx appctx.AppContext) {
	router.Use(middleware.Authentication(config.Env.SecretAccessKey))

	router.Post("/", guestBookHttp.CreateMessage(actx))
	router.Get("/admin", middleware.Authorization(actx, config.Env.SecretAccessKey), guestBookHttp.ListGuestBookForAdmin(actx))
	router.Get("/", guestBookHttp.ListMessage(actx))
	router.Put("/:id", middleware.Authorization(actx, config.Env.SecretAccessKey), guestBookHttp.UpdateMessage(actx))
}
