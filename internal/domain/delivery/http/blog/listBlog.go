package blogHttp

import (
	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	blogBusiness "nta-blog/internal/domain/business/blog"
	blogService "nta-blog/internal/domain/service/blog"
	blogStorage "nta-blog/internal/domain/storage/blog"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func ListBlog(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		logger := apctx.GetLogger()
		mongodb := apctx.GetMongoDB()
		blogType := c.Query("type", string(cnst.BlogTypeConstant.Post))

		store := blogStorage.NewStore(mongodb)
		service := blogService.NewListBlogStore(store)
		biz := blogBusiness.NewListBlogBiz(service)
		result, err := biz.ListBlog(c.Context(), cnst.IBlogType(blogType))
		if err != nil {
			logger.Err(err).Msg("Failed to get list blog")
			panic(err)
		}
		return c.Status(fiber.StatusOK).JSON(common.SimpleSuccessResponse(result))
	}
}
