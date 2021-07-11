package v1

type OutputCallback interface {
	Output(format string, a ...interface{})
	Error(format string, a ...interface{})
	Header(format string, a ...interface{})
	Result(result Result)
}
