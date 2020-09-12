package syserror


type Error404 struct{
	UnknowError
}

func (Error404) Error() string {
	return "非法访问"
}

func (Error404) Code() int {
	return 1002
}

