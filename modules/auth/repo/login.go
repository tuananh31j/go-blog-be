package authrepo

import (
	"context"

	authmdl "nta-blog/modules/auth/model"
)

type LoginStore interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*authmdl.User, error)
	SaveRefreshToken(ctx context.Context, token, userId string) error
}

type loginRepo struct {
	store LoginStore
}

func NewLoginRepo(store LoginStore) *loginRepo {
	return &loginRepo{store: store}
}

func (repo *loginRepo) GetUserByEmail(ctx context.Context, email string) (*authmdl.User, error) {
	return repo.store.FindOneUser(ctx, map[string]interface{}{
		"email": email,
	})
}

func (repo *loginRepo) SaveRefreshToken(ctx context.Context, token, userId string) error {
	return repo.store.SaveRefreshToken(ctx, token, userId)
}
