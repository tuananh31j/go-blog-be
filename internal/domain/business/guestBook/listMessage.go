package guestbookBusiness

import (
	"context"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	guestbookModel "nta-blog/internal/domain/model/guestBook"

	"go.mongodb.org/mongo-driver/bson"
)

type ListMessageService interface {
	GetListMessage(ctx context.Context, pipeline bson.A) ([]guestbookModel.GuestBook, error)
}

type listGuestBookBiz struct {
	service ListMessageService
}

func NewListGuestBookBiz(sv ListMessageService) *listGuestBookBiz {
	return &listGuestBookBiz{
		service: sv,
	}
}

func (biz *listGuestBookBiz) GetListMessage(ctx context.Context, paging, limit uint32) ([]guestbookModel.GuestBook, error) {
	pipeline := bson.A{}

	pipeline = append(pipeline, bson.M{"$match": bson.M{"status": cnst.StatusMessage.Actived}})

	pipeline = append(pipeline, bson.M{"$sort": bson.M{"created_at": -1}, "$skip": paging, "$limit": limit})

	result, err := biz.service.GetListMessage(ctx, pipeline)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return result, nil
}
