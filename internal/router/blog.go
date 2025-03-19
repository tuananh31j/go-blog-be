package routes

import (
	blogHttp "nta-blog/internal/domain/delivery/http/blog"
	"nta-blog/internal/infrastructure/config"
	"nta-blog/internal/lib/appctx"
	"nta-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func blogRouter(router fiber.Router, actx appctx.AppContext) {
	router.Post("/", middleware.Authentication(config.Env.SecretAccessKey), middleware.Authorization(actx, config.Env.SecretAccessKey), blogHttp.CreateBlog(actx))
	router.Get("/admin", middleware.Authentication(config.Env.SecretAccessKey), middleware.Authorization(actx, config.Env.SecretAccessKey), blogHttp.ListBlogForAdmin(actx))
	router.Get("/", blogHttp.ListBlog(actx))
	router.Get("/:id", blogHttp.DetailsBlog(actx))
	// router.Put("/:id", blogHttp.UpdateBlog(actx))
	// router.Delete("/:id", blogHttp.DeleteBlog(actx))
}
