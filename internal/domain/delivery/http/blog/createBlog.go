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

func CreateBlog(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		mongoDb := apctx.GetMongoDB()
		logger := apctx.GetLogger()
		rdb := apctx.GetRedis()
		var payload blogModel.CreatePayload
		err := c.BodyParser(&payload)
		if err != nil {
			logger.Debug().Err(err).Msg("Payload is not valid!")
			panic(err)
		}
		userId, ok := c.Locals("userId").(primitive.ObjectID)
		if !ok {
			logger.Debug().Msg("Failed to get userId from context")
			panic(common.ErrBadRequest(nil))
		}
		payload.UserId = userId
		blogStore := blogStorage.NewStore(mongoDb)
		tagStore := tagStorage.NewStore(mongoDb)
		userStore := userStorage.NewStore(mongoDb, rdb)

		serviceBlog := service.NewBlogService(tagStore, userStore, blogStore)
		biz := blogBusiness.NewCreateBlogBiz(serviceBlog)
		err = biz.CreateBlog(c.Context(), &payload)
		if err != nil {
			logger.Err(err).Msg("Failed to create blog")
			panic(err)
		}
		return c.Status(fiber.StatusCreated).JSON(common.SimpleSuccessResponse("Blog created successfully"))
	}
}
