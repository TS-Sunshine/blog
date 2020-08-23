package syserror

type Error interface {
	Error() string
	Code() int
	ReasonError() error
}
