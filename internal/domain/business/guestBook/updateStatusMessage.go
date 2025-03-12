package guestbookBusiness

import (
	"context"

	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateStatusMessage interface {
	CheckMessageExists(ctx context.Context, msgId primitive.ObjectID) error
	UpdateStatusMessage(ctx context.Context, msgId primitive.ObjectID, status cnst.TStatusMessage) error
}

type updateStatusBiz struct {
	service UpdateStatusMessage
}

func NewUpdateStatusBiz(sv UpdateStatusMessage) *updateStatusBiz {
	return &updateStatusBiz{service: sv}
}

func (biz *updateStatusBiz) UpdateStatus(ctx context.Context, msgId primitive.ObjectID, status cnst.TStatusMessage) error {
	if err := biz.service.CheckMessageExists(ctx, msgId); err != nil {
		return common.ErrBadRequest(err)
	}
	if err := biz.service.UpdateStatusMessage(ctx, msgId, status); err != nil {
		return common.ErrBadRequest(err)
	}
	return nil
}
