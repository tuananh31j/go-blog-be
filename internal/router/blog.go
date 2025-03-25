package routes

import (
	blogHttp "nta-blog/internal/domain/delivery/http/blog"
	"nta-blog/internal/infrastructure/config"
	"nta-blog/internal/lib/appctx"
	"nta-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func blogRouter(router fiber.Router, actx appctx.AppContext) {
	router.Post("/post", middleware.Authentication(config.Env.SecretAccessKey), middleware.Authorization(actx, config.Env.SecretAccessKey), blogHttp.CreateBlog(actx))
	router.Post("/project", middleware.Authentication(config.Env.SecretAccessKey), middleware.Authorization(actx, config.Env.SecretAccessKey), blogHttp.CreateProject(actx))
	router.Post("/me", middleware.Authentication(config.Env.SecretAccessKey), middleware.Authorization(actx, config.Env.SecretAccessKey), blogHttp.MutateMe(actx))
	router.Get("/admin", middleware.Authentication(config.Env.SecretAccessKey), middleware.Authorization(actx, config.Env.SecretAccessKey), blogHttp.ListBlogForAdmin(actx))
	router.Get("/admin/:id", middleware.Authentication(config.Env.SecretAccessKey), middleware.Authorization(actx, config.Env.SecretAccessKey), blogHttp.ForAdmin(actx))
	router.Put("/admin/upload/:id", middleware.Authentication(config.Env.SecretAccessKey), middleware.Authorization(actx, config.Env.SecretAccessKey), blogHttp.UpdateBlog(actx))

	router.Get("/", blogHttp.ListBlog(actx))
	router.Get("/talab", blogHttp.IamtuananhInfo(actx))
	router.Get("/:id", blogHttp.DetailsBlog(actx))
}
