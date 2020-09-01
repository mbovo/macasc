package v1

import (
  "reflect"
  "testing"
)

func TestPrintVars(t *testing.T) {
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
      if got := PrintVars(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("PrintVars() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestSetVars(t *testing.T) {
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
      if got := SetVars(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("SetVars() = %v, want %v", got, tt.want)
      }
    })
  }
}
