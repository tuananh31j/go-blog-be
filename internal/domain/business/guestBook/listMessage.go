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
	TotalDocs(ctx context.Context, pipeline bson.A) (uint32, error)
}

type listGuestBookBiz struct {
	service ListMessageService
}

func NewListGuestBookBiz(sv ListMessageService) *listGuestBookBiz {
	return &listGuestBookBiz{
		service: sv,
	}
}

func (biz *listGuestBookBiz) GetListMessage(ctx context.Context, paging, limit uint32) ([]guestbookModel.GuestBook, uint32, error) {
	pipeline := bson.A{}

	pipeline = append(pipeline, bson.M{"$match": bson.M{"status": cnst.StatusMessage.Actived}})

	pipeline = append(pipeline, bson.M{"$sort": bson.M{"created_at": -1}})
	total, err := biz.service.TotalDocs(ctx, pipeline)
	if err != nil {
		return nil, 0, common.ErrInternal(err)
	}

	pipeline = append(pipeline, bson.M{"$skip": (paging - 1) * limit})
	pipeline = append(pipeline, bson.M{"$limit": limit})

	result, err := biz.service.GetListMessage(ctx, pipeline)
	if err != nil {
		return nil, 0, common.ErrInternal(err)
	}

	return result, total, nil
}
