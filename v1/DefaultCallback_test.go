package v1

import (
  "os"
  "reflect"
  "testing"
)

func TestDefaultCallback_Error(t *testing.T) {
  type fields struct {
    stdout *os.File
    stderr *os.File
  }
  type args struct {
    format string
    a      []interface{}
  }
  tests := []struct {
    name   string
    fields fields
    args   args
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {})
  }
}

func TestDefaultCallback_Header(t *testing.T) {
  type fields struct {
    stdout *os.File
    stderr *os.File
  }
  type args struct {
    format string
    a      []interface{}
  }
  tests := []struct {
    name   string
    fields fields
    args   args
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {})
  }
}

func TestDefaultCallback_Output(t *testing.T) {
  type fields struct {
    stdout *os.File
    stderr *os.File
  }
  type args struct {
    format string
    a      []interface{}
  }
  tests := []struct {
    name   string
    fields fields
    args   args
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {})
  }
}

func TestDefaultCallback_Result(t *testing.T) {
  type fields struct {
    stdout *os.File
    stderr *os.File
  }
  type args struct {
    r Result
  }
  tests := []struct {
    name   string
    fields fields
    args   args
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {})
  }
}

func TestNewDefaultCallback(t *testing.T) {
  type args struct {
    out *os.File
    err *os.File
  }
  tests := []struct {
    name string
    args args
    want OutputCallback
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := NewDefaultCallback(tt.args.out, tt.args.err); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("NewDefaultCallback() = %v, want %v", got, tt.want)
      }
    })
  }
}