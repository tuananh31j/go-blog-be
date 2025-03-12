package userModel

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GoogleLoginDTO struct {
	Id            string `json:"id" bson:"id"`
	Email         string `json:"email" bson:"email"`
	VerifiedEmail bool   `json:"verified_email" bson:"verified_email"`
	Name          string `json:"name" bson:"name"`
	GivenName     string `json:"given_name" bson:"given_name"`
	FamilyName    string `json:"family_name" bson:"family_name"`
	Picture       string `json:"picture" bson:"picture"`
}
