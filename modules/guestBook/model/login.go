package authmdl

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// {
// 	"id": "105574576094292531222",
// 	"email": "tuananhcbqqo@gmail.com",
// 	"verified_email": true,
// 	"name": "Anh Tuấn",
// 	"given_name": "Anh",
// 	"family_name": "Tuấn",
// 	"picture": "https://lh3.googleusercontent.com/a/ACg8ocLz0rZlbNEjy90NsiwSMEvBV9enJXpidAnHC81IdZUvNmc8Dg=s96-c"
//   }

type GoogleLoginDTO struct {
	Id            string `json:"id" bson:"id"`
	Email         string `json:"email" bson:"email"`
	VerifiedEmail bool   `json:"verified_email" bson:"verified_email"`
	Name          string `json:"name" bson:"name"`
	GivenName     string `json:"given_name" bson:"given_name"`
	FamilyName    string `json:"family_name" bson:"family_name"`
	Picture       string `json:"picture" bson:"picture"`
}
