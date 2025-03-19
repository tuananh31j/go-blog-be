package guestbookModel

import (
	"nta-blog/internal/common"
	cnst "nta-blog/internal/constant"
	userModel "nta-blog/internal/domain/model/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const GuestBookCollection = "guest_books"

type GuestBook struct {
	common.CommonModal `bson:",inline"`
	Message            string                 `bson:"message" json:"message"`
	UserId             primitive.ObjectID     `bson:"user_id" json:"user_id"`
	Status             cnst.TStatusMessage    `bson:"status" json:"status"`
	User               *userModel.PrivateUser `bson:",omitempty json:user,omitempty" json:"user,omitempty"`
}
