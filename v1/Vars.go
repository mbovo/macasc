package v1

import (
	"gopkg.in/yaml.v2"

	"github.com/mbovo/yacasc/v1/internal"
)

// Map of variables
type Vars map[string]interface{}

// Load vars from file, can be local or remote (http/https)
func LoadVarFile(uri string) (Vars, error) {
	v := Vars{}
	if data, e := internal.LoadFromFileOrUri(uri); e == nil {
		e = yaml.Unmarshal(data, v)
		if e != nil {
			return v, e
		}
	}
	return v, nil
}
