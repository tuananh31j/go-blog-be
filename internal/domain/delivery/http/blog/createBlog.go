package blogHttp

import (
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func CreateBlog(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
