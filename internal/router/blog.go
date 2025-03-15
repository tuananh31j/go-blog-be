package routes

import (
	blogHttp "nta-blog/internal/domain/delivery/http/blog"
	"nta-blog/internal/infrastructure/config"
	"nta-blog/internal/lib/appctx"
	"nta-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func blogRouter(router fiber.Router, actx appctx.AppContext) {
	router.Use(middleware.Authentication(config.Env.SecretAccessKey))
	router.Use(middleware.Authorization(actx, config.Env.SecretAccessKey))

	router.Post("/", blogHttp.CreateBlog(actx))
	// router.Get("/", blogHttp.ListBlogs(actx))
	// router.Get("/:id", blogHttp.GetBlogByID(actx))
	// router.Put("/:id", blogHttp.UpdateBlog(actx))
	// router.Delete("/:id", blogHttp.DeleteBlog(actx))
}
