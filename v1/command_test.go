package v1

import (
  "reflect"
  "testing"
)

func TestCommand_Execute(t *testing.T) {
  type fields struct {
    Name     string
    Aliases  []string
    Run      func(*Command) Result
    Help     string
    Args     map[string]string
    vars     map[string]interface{}
    args     args
    callback OutputCallback
  }
  type args struct {
    args        map[string]interface{}
    vars        map[string]interface{}
    outCallback OutputCallback
  }
  tests := []struct {
    name   string
    fields fields
    args   args
    want   Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := Command{
        Name:     tt.fields.Name,
        Aliases:  tt.fields.Aliases,
        Run:      tt.fields.Run,
        Help:     tt.fields.Help,
        Args:     tt.fields.Args,
        vars:     tt.fields.vars,
        args:     tt.fields.args,
        callback: tt.fields.callback,
      }
      if got := c.Execute(tt.args.args, tt.args.vars, tt.args.outCallback); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Execute() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestCommand_Is(t *testing.T) {
  type fields struct {
    Name     string
    Aliases  []string
    Run      func(*Command) Result
    Help     string
    Args     map[string]string
    vars     map[string]interface{}
    args     args
    callback OutputCallback
  }
  type args struct {
    name string
  }
  tests := []struct {
    name   string
    fields fields
    args   args
    want   bool
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := Command{
        Name:     tt.fields.Name,
        Aliases:  tt.fields.Aliases,
        Run:      tt.fields.Run,
        Help:     tt.fields.Help,
        Args:     tt.fields.Args,
        vars:     tt.fields.vars,
        args:     tt.fields.args,
        callback: tt.fields.callback,
      }
      if got := c.Is(tt.args.name); got != tt.want {
        t.Errorf("Is() = %v, want %v", got, tt.want)
      }
    })
  }
}