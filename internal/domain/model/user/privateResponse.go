package userModel

import cnst "nta-blog/internal/constant"

type PrivateUser struct {
	Id       string            `json:"id" bson:"_id"`
	NameFake string            `json:"name" bson:"name_fake"`
	Email    string            `json:"email" bson:"email"`
	Role     cnst.TRoleAccount `json:"role" bson:"role"`
	Avt      string            `json:"avt" bson:"avt"`
}
