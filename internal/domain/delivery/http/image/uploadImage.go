package imageHttp

import (
	"nta-blog/internal/common"
	imageBusiness "nta-blog/internal/domain/business/image"
	"nta-blog/internal/domain/service"
	imageStorage "nta-blog/internal/domain/storage/image"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		cld := apctx.GetCloudinary()
		mongoDB := apctx.GetMongoDB()
		logger := apctx.GetLogger()
		fileHeader, err := c.FormFile("file")
		if err != nil {
			logger.Debug().Err(err).Msg("Failed to get file from form")
			panic(err)
		}

		store := imageStorage.NewStore(mongoDB, cld)
		service := service.NewImageService(store)
		biz := imageBusiness.NewUploadImageBiz(service)

		imageData, err := biz.UploadImage(c.Context(), fileHeader)
		if err != nil {
			logger.Debug().Err(err).Msg("Failed to upload image")
			panic(err)
		}

		return c.Status(fiber.StatusCreated).JSON(common.SimpleSuccessResponse(imageData))
	}
}
