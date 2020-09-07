package v2

type CommandResult interface {
  Message() string
  Info() map[string]string
  Status() string
  Error() string
}

const (
  ResultOk      string = "OK"
  ResultError   string = "ERROR"
  ResultSkipped string = "SKIPPED"
  ResultChanged string = "CHANGED"
)

type commandResult struct {
  status  string
  message string
  info    map[string]string
}

func (c commandResult) Message() string {
  return c.message
}

func (c commandResult) Info() map[string]string {
  return c.info
}

func (c commandResult) Status() string {
  return c.status
}

func (c commandResult) Error() string {
  if c.status == ResultError {
    return c.message
  }
  return ""
}

func ReturnOk(message string, info map[string]string) CommandResult {
  return commandResult{
    status:  ResultOk,
    message: message,
    info:    info,
  }
}

func ReturnError(message string, info map[string]string) CommandResult{
  return commandResult{
    status: ResultError,
    message: message,
    info: info,
  }
}

func ReturnSkipped(message string, info map[string]string) CommandResult{
  return commandResult{
    status: ResultSkipped,
    message: message,
    info: info,
  }
}

func ReturnChanged(message string, info map[string]string) CommandResult{
  return commandResult{
    status: ResultChanged,
    message: message,
    info: info,
  }
}