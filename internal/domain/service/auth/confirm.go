package authService

import (
	"context"

	userModel "nta-blog/internal/domain/model/user"
)

type ConfirmStore interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) error
}

type ConfirmTokenStore interface {
	FindToken(ctx context.Context, userId string) (string, error)
}

type confirmService struct {
	userstore  ConfirmStore
	tokenStore ConfirmTokenStore
}

func NewConfirmService(us ConfirmStore, tks ConfirmTokenStore) *confirmService {
	return &confirmService{userstore: us, tokenStore: tks}
}

func (sv *confirmService) FindOneUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error) {
	return sv.userstore.FindOneUser(
		ctx, conditions,
	)
}

func (sv *confirmService) UpdateUser(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) error {
	return sv.userstore.UpdateUser(ctx, filter, data)
}

func (sv *confirmService) GetToken(ctx context.Context, userId string) (string, error) {
	return sv.tokenStore.FindToken(ctx, userId)
}
