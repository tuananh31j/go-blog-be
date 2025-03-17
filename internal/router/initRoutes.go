package routes

import (
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, actx appctx.AppContext) {
	routerV1 := app.Group("/api/v1")

	auth := routerV1.Group("/auth")
	image := routerV1.Group("/images")
	guestBook := routerV1.Group("/guestbook")
	blog := routerV1.Group("/blogs")
	tag := routerV1.Group("/tags")
	user := routerV1.Group("/users")

	imageRouter(image, actx)
	authRoutes(auth, actx)
	guestBookRoutes(guestBook, actx)
	blogRouter(blog, actx)
	tagRoutes(tag, actx)
	userRoutes(user, actx)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong!")
	})
}
