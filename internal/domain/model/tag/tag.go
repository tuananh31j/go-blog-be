package tagModel

import "nta-blog/internal/common"

const TagCollection = "tags"

type Tag struct {
	common.CommonModal `bson:",inline"`
	Name               string `bson:"name" json:"name"`
}
