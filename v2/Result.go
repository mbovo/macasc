package v2

import (
  "fmt"
  "strings"

  "github.com/mbovo/yacasc/v2/command"
)

const (
  OK      string = "OK"
  ERROR   string = "ERROR"
  SKIPPED string = "SKIPPED"
  CHANGED string = "CHANGED"
)

type result struct {
  status  string
  message string
  info    map[string]string
}

func (c result) IsOk() bool {
  return c.status == OK
}

func (c result) IsErr() bool {
  return c.status == ERROR
}

func (c result) IsZero() bool {
  return c.status == ""
}

func (c result) String() string {

  if c.IsZero() { return "" }

  b := strings.Builder{}
  b.WriteString(fmt.Sprintf("[%s]: %s", c.status, c.message))
  for k, v := range c.info{
    b.WriteString(fmt.Sprintf("\n%s: %s", k,v))
  }
  return b.String()
}

func (c result) Message() string {
  return c.message
}

func (c result) Info() map[string]string {
  return c.info
}

func (c result) Status() string {
  return c.status
}

func (c result) Error() string {
  if c.status == ERROR {
    return c.message
  }
  return ""
}

func (c result) HasChanged() bool {
  return c.status == CHANGED
}

func (c result) HasSkipped() bool {
  return c.status == SKIPPED
}

func NewResult(status string, message string, info map[string]string) command.Result {
  return result{
    status: status,
    message: message,
    info: info,
  }
}

func ReturnOk(message string, info map[string]string) command.Result {
  return NewResult(OK, message, info)
}

func ReturnError(message string, info map[string]string) command.Result {
  return NewResult(ERROR, message, info)
}

func ReturnSkipped(message string, info map[string]string) command.Result {
  return NewResult(SKIPPED, message, info)
}

func ReturnChanged(message string, info map[string]string) command.Result {
  return NewResult(CHANGED, message, info)
}