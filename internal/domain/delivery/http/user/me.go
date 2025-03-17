package userHttp

import (
	"nta-blog/internal/common"
	userBusiness "nta-blog/internal/domain/business/user"
	userService "nta-blog/internal/domain/service/user"
	userStorage "nta-blog/internal/domain/storage/user"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMe(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		rdb := apctx.GetRedis()
		mongoDB := apctx.GetMongoDB()
		logger := apctx.GetLogger()

		userId, ok := c.Locals("userId").(primitive.ObjectID)
		if !ok {
			logger.Debug().Msg("Failed to get userId from context")
			panic(common.ErrBadRequest(nil))
		}

		store := userStorage.NewStore(mongoDB, rdb)
		service := userService.NewMeService(store)
		biz := userBusiness.NewMeBiz(service)
		user, err := biz.GetMe(c.Context(), userId)
		if err != nil {

			logger.Debug().Err(err).Msg("Failed to get user")
			panic(err)
		}
		return c.Status(fiber.StatusOK).JSON(common.SimpleSuccessResponse(user))
	}
}
