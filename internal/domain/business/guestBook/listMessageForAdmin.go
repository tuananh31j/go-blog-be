package guestbookBusiness

import (
	"context"
	"strconv"

	"nta-blog/internal/common"
	guestbookModel "nta-blog/internal/domain/model/guestBook"

	"go.mongodb.org/mongo-driver/bson"
)

type ListMessageForAdminService interface {
	GetListMessage(ctx context.Context, pipeline bson.A) ([]guestbookModel.GuestBook, error)
	TotalDocs(ctx context.Context, pipeline bson.A) (uint32, error)
}

type listGuestBookForAdminBiz struct {
	service ListMessageService
}

func NewListGuestBookForAdminBiz(sv ListMessageService) *listGuestBookForAdminBiz {
	return &listGuestBookForAdminBiz{
		service: sv,
	}
}

func (biz *listGuestBookForAdminBiz) GetListMessageForAdmin(ctx context.Context, paging, limit uint32, status string) ([]guestbookModel.GuestBook, uint32, error) {
	pipeline := bson.A{}
	if status != "" {
		state, err := strconv.Atoi(status)
		if err == nil {
			pipeline = append(pipeline, bson.M{"$match": bson.M{"status": state}})
		}
	}
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
	if result == nil {
		return []guestbookModel.GuestBook{}, 0, nil
	}

	return result, total, nil
}
