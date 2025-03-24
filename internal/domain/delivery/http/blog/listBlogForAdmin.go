package blogHttp

import (
	"nta-blog/internal/common"
	blogBusiness "nta-blog/internal/domain/business/blog"
	"nta-blog/internal/domain/service"
	blogStorage "nta-blog/internal/domain/storage/blog"
	tagStorage "nta-blog/internal/domain/storage/tag"
	userStorage "nta-blog/internal/domain/storage/user"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func ListBlogForAdmin(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		logger := apctx.GetLogger()
		mongodb := apctx.GetMongoDB()
		rdb := apctx.GetRedis()

		blogStore := blogStorage.NewStore(mongodb)
		tagStore := tagStorage.NewStore(mongodb)
		userStore := userStorage.NewStore(mongodb, rdb)

		serviceBlog := service.NewBlogService(tagStore, userStore, blogStore)
		biz := blogBusiness.NewListBlogBiz(serviceBlog)
		result, err := biz.ListBlogForAdmin(c.Context())
		if err != nil {
			logger.Err(err).Msg("Failed to get list blog")
			panic(err)
		}
		return c.Status(fiber.StatusOK).JSON(common.SimpleSuccessResponse(result))
	}
}
