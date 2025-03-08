package authBusiness

import (
	"context"

	"nta-blog/common"
	authmdl "nta-blog/modules/auth/model"
	authrepo "nta-blog/modules/auth/repo"
)

type LoginRepo interface {
	authrepo.LoginStore
}

type Hasher interface {
	Hash() string
	SetSalt(s string)
}

type loginRepo struct {
	store  LoginRepo
	hasher Hasher
}

func NewLogin(store LoginRepo) *loginRepo {
	return &loginRepo{store: store}
}

func (biz *loginRepo) Login(ctx context.Context, data *authmdl.LoginDTO) error {
	user, err := biz.store.FindOneUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return common.NewErrorResponse(err, "Not valid!", err.Error())
	}
	biz.hasher.SetSalt(user.Salt)
	hash := biz.hasher.Hash()
	if hash == user.Password {
		return common.NewCustomError(err, "Your data is not valid!", "Password is wrong!")
	}

	return nil
}
