package guestbookBusiness

import (
	"context"
	"strconv"

	"nta-blog/internal/common"
	guestbookModel "nta-blog/internal/domain/model/guestBook"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateStatusMessage interface {
	FindMessage(ctx context.Context, condition map[string]interface{}) ([]*guestbookModel.GuestBook, error)
	UpdateMessage(ctx context.Context, guestBookId primitive.ObjectID, updateField map[string]interface{}) error
}

type updateStatusBiz struct {
	service UpdateStatusMessage
}

func NewUpdateStatusBiz(sv UpdateStatusMessage) *updateStatusBiz {
	return &updateStatusBiz{service: sv}
}

func (biz *updateStatusBiz) UpdateStatus(ctx context.Context, msgId primitive.ObjectID, status string) error {
	state, err := strconv.Atoi(status)
	if err != nil {
		return common.ErrInvalidRequest(err)
	}

	conditions := map[string]interface{}{"_id": msgId}

	updateField := map[string]interface{}{"status": state}
	if _, err := biz.service.FindMessage(ctx, conditions); err != nil {
		return common.ErrBadRequest(err)
	}
	if err := biz.service.UpdateMessage(ctx, msgId, updateField); err != nil {
		return common.ErrBadRequest(err)
	}
	return nil
}
