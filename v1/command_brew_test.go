package v1

import (
  "reflect"
  "testing"
)

func TestBrew(t *testing.T) {
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
      if got := Brew(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Brew() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestBrewAction(t *testing.T) {
  type args struct {
    command  string
    list     []string
    callback OutputCallback
  }
  tests := []struct {
    name    string
    args    args
    wantErr bool
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if err := BrewAction(tt.args.command, tt.args.list, tt.args.callback); (err != nil) != tt.wantErr {
        t.Errorf("BrewAction() error = %v, wantErr %v", err, tt.wantErr)
      }
    })
  }
}
