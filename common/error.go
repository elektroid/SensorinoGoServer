package common

type ErrorType int

const (
	X ErrorType = iota
	SensorinoNotFound
	ServiceNotFound
	ChannelNotFound
)

type Error struct {
	s    string
	Type ErrorType
}

func (e *Error) Error() string {
	return e.s
}

func ConvertError(e error, t ErrorType) *Error {
	if e == nil {
		return nil
	}
	return &Error{e.Error(), t}
}

func NewError(msg string, t ErrorType) *Error {
	return &Error{msg, t}
}
