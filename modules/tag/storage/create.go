package blogstrg

import (
	"context"

	blogmdl "nta-blog/modules/blog/model"
)

func (s *splStore) CreateBlog(ctx context.Context, data *blogmdl.CreateBlog) error {
	s.db.Create(data)
	return nil
}
