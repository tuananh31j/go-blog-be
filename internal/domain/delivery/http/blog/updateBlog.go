package blogHttp

import (
	"nta-blog/internal/common"
	blogBusiness "nta-blog/internal/domain/business/blog"
	blogModel "nta-blog/internal/domain/model/blog"
	"nta-blog/internal/domain/service"
	blogStorage "nta-blog/internal/domain/storage/blog"
	tagStorage "nta-blog/internal/domain/storage/tag"
	userStorage "nta-blog/internal/domain/storage/user"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateBlog(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		logger := apctx.GetLogger()
		mongodb := apctx.GetMongoDB()
		rdb := apctx.GetRedis()
		id := c.Params("id")
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.Debug().Err(err).Msg("Failed to convert id to object id")
			panic(err)
		}
		var payload blogModel.UpdatePayload
		err = c.BodyParser(&payload)
		if err != nil {
			logger.Debug().Err(err).Msg("Payload is not valid!")
			panic(err)
		}

		blogStore := blogStorage.NewStore(mongodb)
		tagStore := tagStorage.NewStore(mongodb)
		userStore := userStorage.NewStore(mongodb, rdb)

		serviceBlog := service.NewBlogService(tagStore, userStore, blogStore)

		biz := blogBusiness.NewUpdateBiz(serviceBlog)
		err = biz.Update(c.Context(), objId, &payload)
		if err != nil {
			logger.Err(err).Msg("Failed to get info")
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.SimpleSuccessResponse("Update successfully"))
	}
}
