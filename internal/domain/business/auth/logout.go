package authBusiness

import "context"

type LogoutService interface {
	RemoveRefreshToken(ctx context.Context, userId string) error
}

type logoutBiz struct {
	service LogoutService
}

func NewLogoutBiz(sv LogoutService) *logoutBiz {
	return &logoutBiz{service: sv}
}

func (biz *logoutBiz) Logout(ctx context.Context, userId string) error {
	if err := biz.service.RemoveRefreshToken(ctx, userId); err != nil {
		return err
	}
	return nil
}
