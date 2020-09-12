package syserror

// 接口
type Error interface {
	Error() string
	Code() int
	ReasonError() error
}

func NewError(msg string, reason error) Error {
	return UnknowError{msg: msg, reason: reason}
}