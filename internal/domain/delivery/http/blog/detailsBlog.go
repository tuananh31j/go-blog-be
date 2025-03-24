package blogHttp

import (
	"strconv"

	"nta-blog/internal/common"
	blogBusiness "nta-blog/internal/domain/business/blog"
	"nta-blog/internal/domain/service"
	blogStorage "nta-blog/internal/domain/storage/blog"
	tagStorage "nta-blog/internal/domain/storage/tag"
	userStorage "nta-blog/internal/domain/storage/user"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DetailsBlog(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		logger := apctx.GetLogger()
		mongodb := apctx.GetMongoDB()
		rdb := apctx.GetRedis()
		isForMetadata, err := strconv.ParseBool(c.Query("metadata", "false"))
		if err != nil {
			logger.Debug().Err(err).Msg("Failed to parse metadata query parameter")
			isForMetadata = false
		}
		id := c.Params("id")
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.Debug().Err(err).Msg("Failed to convert id to object id")
			panic(err)
		}

		blogStore := blogStorage.NewStore(mongodb)
		tagStore := tagStorage.NewStore(mongodb)
		userStore := userStorage.NewStore(mongodb, rdb)

		serviceBlog := service.NewBlogService(tagStore, userStore, blogStore)

		biz := blogBusiness.NewDetailsBlogBiz(serviceBlog)
		result, err := biz.FindDetailsBlog(c.Context(), objId, isForMetadata)
		if err != nil {
			logger.Err(err).Msg("Failed to get details blog")
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.SimpleSuccessResponse(result))
	}
}
