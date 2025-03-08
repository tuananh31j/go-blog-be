package blogstrg

import (
	"context"

	blogmdl "nta-blog/modules/blog/model"
)

func (s *splStore) FindBlogDetails(c context.Context, conditions map[string]interface{}, moreKeys ...string) (*blogmdl.CreateBlog, error) {
	var data blogmdl.CreateBlog
	if err := s.db.Where(conditions).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
