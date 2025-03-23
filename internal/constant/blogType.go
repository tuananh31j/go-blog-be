package cnst

// type Role struct {
// 	user 0,
// 	admin 1
// }

type IBlogType string

type BlogType struct {
	Post    IBlogType
	Project IBlogType
	Me      IBlogType
}

var BlogTypeConstant = BlogType{
	Post:    "post",
	Project: "project",
	Me:      "me",
}
