package userModel

type RegisterDTO struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
