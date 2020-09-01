package v1

import (
  "reflect"
  "testing"
)

func TestGetArgumentFromVars(t *testing.T) {
  type args struct {
    c *Command
  }
  tests := []struct {
    name      string
    args      args
    wantSlice []string
    wantR     *Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      gotSlice, gotR := GetArgumentFromVars(tt.args.c)
      if !reflect.DeepEqual(gotSlice, tt.wantSlice) {
        t.Errorf("GetArgumentFromVars() gotSlice = %v, want %v", gotSlice, tt.wantSlice)
      }
      if !reflect.DeepEqual(gotR, tt.wantR) {
        t.Errorf("GetArgumentFromVars() gotR = %v, want %v", gotR, tt.wantR)
      }
    })
  }
}

func TestGetBoolArgument(t *testing.T) {
  type args struct {
    c            *Command
    argumentName string
  }
  tests := []struct {
    name  string
    args  args
    wantB bool
    wantR *Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      gotB, gotR := GetBoolArgument(tt.args.c, tt.args.argumentName)
      if gotB != tt.wantB {
        t.Errorf("GetBoolArgument() gotB = %v, want %v", gotB, tt.wantB)
      }
      if !reflect.DeepEqual(gotR, tt.wantR) {
        t.Errorf("GetBoolArgument() gotR = %v, want %v", gotR, tt.wantR)
      }
    })
  }
}

func TestGetIntArgument(t *testing.T) {
  type args struct {
    c            *Command
    argumentName string
  }
  tests := []struct {
    name  string
    args  args
    wantN int
    wantR *Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      gotN, gotR := GetIntArgument(tt.args.c, tt.args.argumentName)
      if gotN != tt.wantN {
        t.Errorf("GetIntArgument() gotN = %v, want %v", gotN, tt.wantN)
      }
      if !reflect.DeepEqual(gotR, tt.wantR) {
        t.Errorf("GetIntArgument() gotR = %v, want %v", gotR, tt.wantR)
      }
    })
  }
}

func TestGetStringArgument(t *testing.T) {
  type args struct {
    c            *Command
    argumentName string
  }
  tests := []struct {
    name  string
    args  args
    wantS string
    wantR *Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      gotS, gotR := GetStringArgument(tt.args.c, tt.args.argumentName)
      if gotS != tt.wantS {
        t.Errorf("GetStringArgument() gotS = %v, want %v", gotS, tt.wantS)
      }
      if !reflect.DeepEqual(gotR, tt.wantR) {
        t.Errorf("GetStringArgument() gotR = %v, want %v", gotR, tt.wantR)
      }
    })
  }
}

func TestGetStringArrayArgument(t *testing.T) {
  type args struct {
    c            *Command
    argumentName string
  }
  tests := []struct {
    name  string
    args  args
    wantL []string
    wantR *Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      gotL, gotR := GetStringArrayArgument(tt.args.c, tt.args.argumentName)
      if !reflect.DeepEqual(gotL, tt.wantL) {
        t.Errorf("GetStringArrayArgument() gotL = %v, want %v", gotL, tt.wantL)
      }
      if !reflect.DeepEqual(gotR, tt.wantR) {
        t.Errorf("GetStringArrayArgument() gotR = %v, want %v", gotR, tt.wantR)
      }
    })
  }
}

func TestVerifyRequiredArgs(t *testing.T) {
  type args struct {
    c        *Command
    required []string
  }
  tests := []struct {
    name string
    args args
    want *Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := VerifyRequiredArgs(tt.args.c, tt.args.required); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("VerifyRequiredArgs() = %v, want %v", got, tt.want)
      }
    })
  }
}

func Test_getArgument(t *testing.T) {
  type args struct {
    c    *Command
    name string
  }
  tests := []struct {
    name  string
    args  args
    want  interface{}
    want1 *Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      got, got1 := getArgument(tt.args.c, tt.args.name)
      if !reflect.DeepEqual(got, tt.want) {
        t.Errorf("getArgument() got = %v, want %v", got, tt.want)
      }
      if !reflect.DeepEqual(got1, tt.want1) {
        t.Errorf("getArgument() got1 = %v, want %v", got1, tt.want1)
      }
    })
  }
}
