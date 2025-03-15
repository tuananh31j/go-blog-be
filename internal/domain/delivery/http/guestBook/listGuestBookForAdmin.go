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

func ListGuestBookForAdmin(apctx appctx.AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		logger := apctx.GetLogger()
		mongodb := apctx.GetMongoDB()
		paging := c.Query("page")
		limit := c.Query("limit")
		status := c.Query("status")
		logger.Debug().Msgf("paging: %s, limit: %s, status: %s", paging, limit, status)

		storeGuestBook := guestBookStorage.NewStore(mongodb)
		serviceGuestBook := guestbookService.NewListGuestBookService(storeGuestBook)
		biz := guestbookBusiness.NewListGuestBookForAdminBiz(serviceGuestBook)

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
		data, total, err := biz.GetListMessageForAdmin(c.Context(), page, length, status)
		if err != nil {
			logger.Err(err).Msg("Failed to get list message")
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(data, page, total, map[string]interface{}{"limit": length, "status": status}))
	}
}
