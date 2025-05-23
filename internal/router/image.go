package routes

import (
	imageHttp "nta-blog/internal/domain/delivery/http/image"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func imageRouter(router fiber.Router, actx appctx.AppContext) {
	// router.Use(middleware.Authentication(config.Env.SecretAccessKey))
	// router.Use(middleware.Authorization(actx, config.Env.SecretAccessKey))

	router.Post("/upload", imageHttp.UploadImage(actx))
	router.Get("/", imageHttp.ListImages(actx))
}
