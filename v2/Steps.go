package v2

import (
  "gopkg.in/yaml.v2"
)

type Step interface {
  Runnable
  yaml.Marshaler
  yaml.Unmarshaler
}

type step struct{}

func (s step) Run(vars VarMap, cb OutCallback) {
  panic("implement me")
}

func (s step) MarshalYAML() (interface{}, error) {
  panic("implement me")
}

func (s step) UnmarshalYAML(unmarshal func(interface{}) error) error {
  panic("implement me")
}
