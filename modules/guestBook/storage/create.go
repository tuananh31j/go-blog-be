package authtore

import (
	"context"

	cnst "nta-blog/constant"
	authmdl "nta-blog/modules/auth/model"
)

func (s *store) Create(ctx context.Context, dto *authmdl.User) error {
	col := s.db.Collection(cnst.UserCollection)

	_, err := col.InsertOne(ctx, &dto)

	return err
}
