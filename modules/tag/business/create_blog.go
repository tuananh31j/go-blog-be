package business

import (
	"context"

	blogmdl "nta-blog/modules/blog/model"
)

type CreateBlogStore interface {
	CreateBlog(ctx context.Context, data *blogmdl.CreateBlog) error
}

type createBlog struct {
	store CreateBlogStore
}

func NewCreateBlog(store CreateBlogStore) *createBlog {
	return &createBlog{store}
}

func (biz *createBlog) CreateBlog(ctx context.Context, data *blogmdl.CreateBlog) error {
	// có logix thì thì làm ở đây không thì cứ gọi store bình thường
	// validate các kiểu
	biz.store.CreateBlog(ctx, data)
	return nil
}
