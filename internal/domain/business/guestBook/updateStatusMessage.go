package guestbookBusiness

import (
	"context"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	guestbookModel "nta-blog/internal/domain/model/guestBook"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateStatusMessage interface {
	CheckMessageExists(ctx context.Context, condition map[string]interface{}) ([]*guestbookModel.GuestBook, error)
	UpdateStatusMessage(ctx context.Context, msgId primitive.ObjectID, status cnst.TStatusMessage) error
}

type updateStatusBiz struct {
	service UpdateStatusMessage
}

func NewUpdateStatusBiz(sv UpdateStatusMessage) *updateStatusBiz {
	return &updateStatusBiz{service: sv}
}

func (biz *updateStatusBiz) UpdateStatus(ctx context.Context, msgId primitive.ObjectID, status cnst.TStatusMessage) error {
	conditions := map[string]interface{}{"_id": msgId}
	if _, err := biz.service.CheckMessageExists(ctx, conditions); err != nil {
		return common.ErrBadRequest(err)
	}
	if err := biz.service.UpdateStatusMessage(ctx, msgId, status); err != nil {
		return common.ErrBadRequest(err)
	}
	return nil
}
