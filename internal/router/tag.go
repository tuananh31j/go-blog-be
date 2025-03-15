package routes

import (
	tagHttp "nta-blog/internal/domain/delivery/http/tag"
	"nta-blog/internal/infrastructure/config"
	"nta-blog/internal/lib/appctx"
	"nta-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func tagRoutes(router fiber.Router, actx appctx.AppContext) {
	router.Post("/", middleware.Authentication(config.Env.SecretAccessKey), middleware.Authorization(actx, config.Env.SecretAccessKey), tagHttp.CreateTag(actx))
	router.Get("/", tagHttp.ListTags(actx))
	// router.Get("/:id", tagHttp.GetTagByID(actx))
	// router.Put("/:id", tagHttp.UpdateTag(actx))
	// router.Delete("/:id", tagHttp.DeleteTag(actx))
}
