package imageHttp

import (
	"nta-blog/internal/common"
	imageBusiness "nta-blog/internal/domain/business/image"
	imageService "nta-blog/internal/domain/service/image"
	imageStorage "nta-blog/internal/domain/storage/image"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func ListImages(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("refreshToken")

		cld := apctx.GetCloudinary()
		mongoDB := apctx.GetMongoDB()
		logger := apctx.GetLogger()
		logger.Debug().Msgf("%v", token)
		store := imageStorage.NewStore(mongoDB, cld)
		service := imageService.NewListImageService(store)
		biz := imageBusiness.NewListImagesBiz(service)
		data, err := biz.GetListImage(c.Context())
		if err != nil {
			logger.Debug().Err(err).Msg("Failed to get list images")
			panic(err)
		}
		return c.Status(fiber.StatusOK).JSON(common.SimpleSuccessResponse(data))
	}
}
