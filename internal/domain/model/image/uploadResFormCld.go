package imageModel

type UploadResFormCld struct {
	PublicID  string `bson:"public_id" json:"public_id"`
	Width     uint16 `bson:"width" json:"width"`
	Height    uint16 `bson:"height" json:"height"`
	URL       string `bson:"url" json:"url"`
	SecureURL string `bson:"secure_url" json:"secure_url"`
}
