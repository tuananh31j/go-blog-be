package tagHttp

import (
	"nta-blog/internal/common"
	tagBusiness "nta-blog/internal/domain/business/tag"
	"nta-blog/internal/domain/service"
	tagStorage "nta-blog/internal/domain/storage/tag"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func ListTags(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		mongoDB := apctx.GetMongoDB()
		logger := apctx.GetLogger()

		tagStore := tagStorage.NewStore(mongoDB)
		service := service.NewTagService(tagStore)
		biz := tagBusiness.NewListTagBiz(service)
		data, err := biz.GetAllTag(c.Context())
		if err != nil {
			logger.Debug().Err(err).Msg("Failed to list tags")
			panic(err)
		}
		return c.Status(fiber.StatusOK).JSON(common.SimpleSuccessResponse(data))
	}
}
