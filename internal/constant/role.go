package cnst

// type Role struct {
// 	user 0,
// 	admin 1
// }

type TRoleAccount uint8

type role struct {
	User  TRoleAccount
	Admin TRoleAccount
}

var Role = role{
	User:  0,
	Admin: 1,
}
