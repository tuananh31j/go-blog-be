package authrepo

import (
	"context"

	authmdl "nta-blog/modules/auth/model"
)

type LoginStore interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*authmdl.User, error)
	SaveRefreshToken(ctx context.Context, token, userId string) error
}

type loginStore struct {
	store LoginStore
}

func NewLoginStore(store LoginStore) *loginStore {
	return &loginStore{store: store}
}

func (repo *loginStore) GetUserByEmail(ctx context.Context, email string) (*authmdl.User, error) {
	return repo.store.FindOneUser(ctx, map[string]interface{}{
		"email": email,
	})
}

func (repo *loginStore) SaveRefreshToken(ctx context.Context, token, userId string) error {
	return repo.store.SaveRefreshToken(ctx, token, userId)
}
