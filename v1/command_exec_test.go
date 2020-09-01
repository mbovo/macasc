package v1

import (
  "reflect"
  "testing"
)

func TestExec(t *testing.T) {
  type args struct {
    c *Command
  }
  tests := []struct {
    name string
    args args
    want Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := Exec(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Exec() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestShell(t *testing.T) {
  type args struct {
    c *Command
  }
  tests := []struct {
    name string
    args args
    want Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := Shell(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Shell() = %v, want %v", got, tt.want)
      }
    })
  }
}
