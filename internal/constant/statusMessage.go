package cnst

type TStatusMessage uint8

type statusMessage struct {
	Pending TStatusMessage
	Actived TStatusMessage
	Hidden  TStatusMessage
}

var StatusMessage = statusMessage{
	Pending: 0,
	Actived: 1,
	Hidden:  2,
}
