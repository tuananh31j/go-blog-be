package authBusiness

// import (
// 	"context"

// 	"nta-blog/common"
// 	authmodel "nta-blog/modules/auth/model"
// )

// type RegisterStore interface {
// 	Register(ctx *context.Context, dto *authmodel.RegisterDTO) error
// }

// type registerStore struct {
// 	store RegisterStore
// }

// func NewRegistor(store RegisterStore) *registerStore {
// 	return &registerStore{store: store}
// }

// func (biz *registerStore) Register(ctx *context.Context, dto *authmodel.RegisterDTO) error {
// 	if err := biz.store.Register(ctx, dto); err != nil {
// 		return common.ErrInternal(err)
// 	}
// 	return nil
// }
