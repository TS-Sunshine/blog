package syserror

type UnknowError interface {
	Error() string
	Code() int
	ReasonError() error
}
