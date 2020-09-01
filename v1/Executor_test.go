package v1

import (
  "reflect"
  "testing"
)

func TestExecutor_Run(t *testing.T) {
  type fields struct {
    callback  OutputCallback
    VarFile   string
    StepsFile string
    Steps     []Step
    Vars      Vars
    Commands  []Command
  }
  type args struct {
    action Action
  }
  tests := []struct {
    name    string
    fields  fields
    args    args
    wantErr bool
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      ex := Executor{
        callback:  tt.fields.callback,
        VarFile:   tt.fields.VarFile,
        StepsFile: tt.fields.StepsFile,
        Steps:     tt.fields.Steps,
        Vars:      tt.fields.Vars,
        Commands:  tt.fields.Commands,
      }
      if err := ex.Run(tt.args.action); (err != nil) != tt.wantErr {
        t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
      }
    })
  }
}

func TestNewExecutor(t *testing.T) {
  type args struct {
    callback  OutputCallback
    StepsFile string
    VarFile   string
  }
  tests := []struct {
    name    string
    args    args
    want    *Executor
    wantErr bool
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      got, err := NewExecutor(tt.args.callback, tt.args.StepsFile, tt.args.VarFile)
      if (err != nil) != tt.wantErr {
        t.Errorf("NewExecutor() error = %v, wantErr %v", err, tt.wantErr)
        return
      }
      if !reflect.DeepEqual(got, tt.want) {
        t.Errorf("NewExecutor() got = %v, want %v", got, tt.want)
      }
    })
  }
}