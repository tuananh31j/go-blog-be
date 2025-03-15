package tagHttp

import (
	"nta-blog/internal/common"
	tagBusiness "nta-blog/internal/domain/business/tag"
	tagModel "nta-blog/internal/domain/model/tag"
	tagService "nta-blog/internal/domain/service/tag"
	tagStorage "nta-blog/internal/domain/storage/tag"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func CreateTag(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		mongodb := apctx.GetMongoDB()
		logger := apctx.GetLogger()
		var payload tagModel.TagDTO
		if err := c.BodyParser(&payload); err != nil {
			logger.Debug().Err(err).Msg("Payload is not valid!")
			panic(common.ErrBadRequest(err))
		}

		tagStore := tagStorage.NewStore(mongodb)
		tagService := tagService.NewCreateTagService(tagStore)
		biz := tagBusiness.NewCreateTagBiz(tagService)
		err := biz.CreateNewTag(c.Context(), &payload)
		if err != nil {
			logger.Debug().Err(err).Msg("Failed to create tag")
			panic(err)
		}

		return c.Status(fiber.StatusCreated).JSON(common.SimpleSuccessResponse("Tag created successfully"))
	}
}
