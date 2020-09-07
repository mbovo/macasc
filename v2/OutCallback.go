package v2

type OutCallback interface {
  Output(format string, a ...interface{})
  Error(format string, a ...interface{})
  Header(format string, a ...interface{})
  Result(result CommandResult)
}