package authmdl

type User struct {
	Name     string
	Password string
	Salt     string
	Email    string
	Role     int
}
