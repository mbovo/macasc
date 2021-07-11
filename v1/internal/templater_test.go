package internal

import (
	"gopkg.in/flosch/pongo2.v3"
	"strings"
	"testing"
)

func TestTemplater(t *testing.T) {
	v := pongo2.Context{
		"name": "yacasc",
		"desc": "Yet Another Configurator As Code",
	}
	s := "{{ name }} {{desc}}"
	result, err := Template(s, v)
	if err != nil {
		t.Errorf("Error templating %v", err)
	}
	if !strings.Contains(result, "yacasc") {
		t.Errorf("Template is wrong!")
	}
}
