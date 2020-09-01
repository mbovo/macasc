package v1

import (
  "reflect"
  "testing"
)

func TestHashString(t *testing.T) {
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
      if got := HashString(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("HashString() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestHashfile(t *testing.T) {
  type args struct {
    c *Command
  }
  tests := []struct {
    name    string
    args    args
    wantRes Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if gotRes := Hashfile(tt.args.c); !reflect.DeepEqual(gotRes, tt.wantRes) {
        t.Errorf("Hashfile() = %v, want %v", gotRes, tt.wantRes)
      }
    })
  }
}
