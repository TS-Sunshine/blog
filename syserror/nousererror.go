package syserror

type NoUserError struct {
	UnknowError
}

func (this NoUserError) Error() string {
	return "请登录系统"
}

func (this NoUserError) Code() int {
	return 1001
}


