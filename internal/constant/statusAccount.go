package cnst

type TStatusAccount uint8

type statusAccount struct {
	Actived TStatusAccount
	Banned  TStatusAccount
}

var StatusAccount = statusAccount{
	Actived: 0,
	Banned:  1,
}
