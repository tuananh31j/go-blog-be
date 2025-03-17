package authHttp

import (
	"nta-blog/internal/common"
	authBusiness "nta-blog/internal/domain/business/auth"
	authService "nta-blog/internal/domain/service/auth"
	tokenStorage "nta-blog/internal/domain/storage/token"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Logout(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		rdb := apctx.GetRedis()
		logger := apctx.GetLogger()
		store := tokenStorage.NewStore(apctx.GetMongoDB(), rdb)
		service := authService.NewLogoutService(store)
		biz := authBusiness.NewLogoutBiz(service)

		userId, ok := c.Locals("userId").(primitive.ObjectID)
		if !ok {
			logger.Debug().Msg("Failed to get userId from context")
			panic(common.ErrBadRequest(nil))
		}
		err := biz.Logout(c.Context(), userId.Hex())
		if err != nil {
			logger.Debug().Msg(err.Error())
			panic(err)
		}
		c.ClearCookie("accessToken")
		c.ClearCookie("refreshToken")

		return c.Status(fiber.StatusOK).JSON(common.SimpleSuccessResponse("Logout successfully"))
	}
}
