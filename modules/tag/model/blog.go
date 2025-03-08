package blogmdl

import "nta-blog/common"

type CreateBlog struct {
	common.SQLModal
	Title   string `json:"title"`
	Content string `json:"content"`
}
