package error

type Error interface {
	Caller() []CallerInfo
	Wrapped() []error
	Code() int
	error
	private()
}

type CallerInfo struct {
	FuncName string
	FileName string
	FileLine int
}
