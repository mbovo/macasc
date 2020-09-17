package command

import (
  "fmt"

  "gopkg.in/yaml.v2"
)

type OutCallback interface {
  Output(format string, a ...interface{})
  Error(format string, a ...interface{})
  Header(format string, a ...interface{})
  Result(result Result)
}

type VarMap interface {
  Add(string, interface{})
  Remove(string)
  Get(string) interface{}
  Template(string) (string, error)
  yaml.Marshaler
  yaml.Unmarshaler
}

type Command interface{
  Name() string
  Exec(vars VarMap) Result
  Help(callback OutCallback)
  Is(name string) bool
  Aliases() []string
}

type Result interface {
  Message() string
  Info() map[string]string
  Status() string
  IsOk() bool
  IsErr() bool
  IsZero() bool
  HasChanged() bool
  HasSkipped() bool
  error
  fmt.Stringer
}