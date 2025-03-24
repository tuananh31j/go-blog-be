package guestBookHttp

import (
	guestbookBusiness "nta-blog/internal/domain/business/guestBook"
	"nta-blog/internal/domain/service"
	guestBookStorage "nta-blog/internal/domain/storage/guestBook"
	userStorage "nta-blog/internal/domain/storage/user"
	"nta-blog/internal/lib/appctx"
	"nta-blog/internal/lib/logger"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlaceHolder struct {
	Status string `json:"status"`
}

func UpdateMessage(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		mongodb := apctx.GetMongoDB()
		rdb := apctx.GetRedis()
		logger := logger.NewLogger()

		id := c.Params("id")
		var payload PlaceHolder

		err := c.BodyParser(&payload)
		if err != nil {
			logger.Debug().Err(err).Msg("Payload is not valid!")
			panic(err)
		}
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.Debug().Err(err).Msg("Failed to convert id to object id")
			panic(err)
		}
		guestBookStore := guestBookStorage.NewStore(mongodb)
		userStore := userStorage.NewStore(mongodb, rdb)
		service := service.NewGuestBookService(guestBookStore, userStore)
		biz := guestbookBusiness.NewUpdateStatusBiz(service)
		biz.UpdateStatus(c.Context(), objId, payload.Status)
		return nil
	}
}
