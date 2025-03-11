package authrepo

import (
	"context"

	authmdl "nta-blog/modules/auth/model"
)

type GoogleLoginStore interface {
	FindOneUser(ctx context.Context, conditions map[string]interface{}) (*authmdl.User, error)
	SaveRefreshToken(ctx context.Context, token, userId string) error
	Create(ctx context.Context, dto *authmdl.User) error
}

type GoogleloginRepo struct {
	store GoogleLoginStore
}

func NewGoogleLoginRepo(store GoogleLoginStore) *GoogleloginRepo {
	return &GoogleloginRepo{store: store}
}

func (repo *GoogleloginRepo) GetUserByEmail(ctx context.Context, email string) (*authmdl.User, error) {
	return repo.store.FindOneUser(ctx, map[string]interface{}{
		"email": email,
	})
}

func (repo *GoogleloginRepo) CreateUser(ctx context.Context, dto *authmdl.User) error {
	return repo.store.Create(ctx, dto)
}

func (repo *GoogleloginRepo) SaveRefreshToken(ctx context.Context, token, userId string) error {
	return repo.store.SaveRefreshToken(ctx, token, userId)
}
