package business

import (
	"context"
	"errors"
	"testing"

	blogmdl "nta-blog/modules/blog/model"

	"github.com/stretchr/testify/assert"
)

type mockCreateBlogStore struct {
	err error
}

func (m *mockCreateBlogStore) CreateBlog(ctx context.Context, data *blogmdl.CreateBlog) error {
	return m.err
}

func TestCreateBlog(t *testing.T) {
	t.Run("successful creation", func(t *testing.T) {
		mockStore := &mockCreateBlogStore{}
		biz := NewCreateBlog(mockStore)

		data := &blogmdl.CreateBlog{
			Title:   "Test Blog",
			Content: "This is a test blog content",
		}

		err := biz.CreateBlog(context.Background(), data)
		assert.NoError(t, err)
	})

	t.Run("store returns error", func(t *testing.T) {
		mockStore := &mockCreateBlogStore{err: errors.New("store error")}
		biz := NewCreateBlog(mockStore)

		data := &blogmdl.CreateBlog{
			Title:   "Test Blog",
			Content: "This is a test blog content",
		}

		err := biz.CreateBlog(context.Background(), data)
		assert.Error(t, err)
		assert.Equal(t, "store error", err.Error())
	})
}
