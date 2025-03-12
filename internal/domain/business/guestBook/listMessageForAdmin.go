package guestbookBusiness

import (
	"context"

	"nta-blog/internal/common"
	guestbookModel "nta-blog/internal/domain/model/guestBook"

	"go.mongodb.org/mongo-driver/bson"
)

type ListMessageForAdminService interface {
	GetListMessage(ctx context.Context, pipeline bson.A) ([]guestbookModel.GuestBook, error)
}

type listGuestBookForAdminBiz struct {
	service ListMessageService
}

func NewListGuestBookForAdminBiz(sv ListMessageService) *listGuestBookForAdminBiz {
	return &listGuestBookForAdminBiz{
		service: sv,
	}
}

func (biz *listGuestBookForAdminBiz) GetListMessageForAdmin(ctx context.Context, paging, limit uint32) ([]guestbookModel.GuestBook, error) {
	pipeline := bson.A{}

	pipeline = append(pipeline, bson.M{"$sort": bson.M{"created_at": -1}, "$skip": paging, "$limit": limit})

	result, err := biz.service.GetListMessage(ctx, pipeline)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return result, nil
}
