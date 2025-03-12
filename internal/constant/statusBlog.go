package cnst

type TStatusBlog uint8

type statusBlog struct {
	Pending TStatusBlog
	Actived TStatusBlog
	Hidden  TStatusBlog
}

var StatusBlog = statusBlog{
	Pending: 0,
	Actived: 1,
	Hidden:  2,
}
