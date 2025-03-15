package guestBookHttp

import (
	"strconv"

	"nta-blog/internal/common"
	guestbookBusiness "nta-blog/internal/domain/business/guestBook"
	guestbookService "nta-blog/internal/domain/service/guestBook"
	guestBookStorage "nta-blog/internal/domain/storage/guestBook"
	"nta-blog/internal/lib/appctx"

	"github.com/gofiber/fiber/v2"
)

func ListMessage(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		logger := apctx.GetLogger()
		mongodb := apctx.GetMongoDB()
		paging := c.Query("page")
		limit := c.Query("limit")

		storeGuestBook := guestBookStorage.NewStore(mongodb)
		serviceGuestBook := guestbookService.NewListGuestBookService(storeGuestBook)
		biz := guestbookBusiness.NewListGuestBookBiz(serviceGuestBook)
		limitInt, err := strconv.Atoi(limit)
		length := uint32(limitInt)
		if err != nil {
			length = 10
		}
		pageInt, err := strconv.Atoi(paging)
		page := uint32(pageInt)
		if err != nil {
			page = 1
		}
		data, total, err := biz.GetListMessage(c.Context(), page, length)
		if err != nil {
			logger.Err(err).Msg("Failed to get list message")
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(data, page, total, map[string]interface{}{"limit": limit}))
	}
}
