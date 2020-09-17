package v2

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


