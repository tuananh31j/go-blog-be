package guestBookHttp

import (
	guestbookBusiness "nta-blog/internal/domain/business/guestBook"
	guestbookService "nta-blog/internal/domain/service/guestBook"
	guestBookStorage "nta-blog/internal/domain/storage/guestBook"
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
		mongoDB := apctx.GetMongoDB()
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
		store := guestBookStorage.NewStore(mongoDB)
		service := guestbookService.NewUpdateGuestBookService(store)
		biz := guestbookBusiness.NewUpdateStatusBiz(service)
		biz.UpdateStatus(c.Context(), objId, payload.Status)
		return nil
	}
}
