package v2

import (
  "gopkg.in/yaml.v2"
)

type VarMap interface {
  Add(string, interface{})
  Remove(string)
  Get(string) interface{}
  Template(string) (string, error)
  yaml.Marshaler
  yaml.Unmarshaler
}

type varmap struct {}

func (v varmap) Add(s string, i interface{}) {
  panic("implement me")
}

func (v varmap) Remove(s string) {
  panic("implement me")
}

func (v varmap) Get(s string) interface{} {
  panic("implement me")
}

func (v varmap) Template(s string) (string, error) {
  panic("implement me")
}

func (v varmap) MarshalYAML() (interface{}, error) {
  panic("implement me")
}

func (v varmap) UnmarshalYAML(unmarshal func(interface{}) error) error {
  panic("implement me")
}


