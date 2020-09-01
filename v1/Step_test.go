package v1

import (
  "reflect"
  "testing"
)

func TestLoadStepFile(t *testing.T) {
  type args struct {
    uri     string
    context map[string]interface{}
  }
  tests := []struct {
    name    string
    args    args
    want    []Step
    wantErr bool
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      got, err := LoadStepFile(tt.args.uri, tt.args.context)
      if (err != nil) != tt.wantErr {
        t.Errorf("LoadStepFile() error = %v, wantErr %v", err, tt.wantErr)
        return
      }
      if !reflect.DeepEqual(got, tt.want) {
        t.Errorf("LoadStepFile() got = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestStep_Run(t *testing.T) {
  type fields struct {
    Name              string
    AddCommands       []map[string]interface{}
    RemoveCommands    []map[string]interface{}
    VerifyCommands    []map[string]interface{}
    ConfigureCommands []map[string]interface{}
  }
  type args struct {
    ex      Executor
    name    string
    cmdList []map[string]interface{}
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
      s := Step{
        Name:              tt.fields.Name,
        AddCommands:       tt.fields.AddCommands,
        RemoveCommands:    tt.fields.RemoveCommands,
        VerifyCommands:    tt.fields.VerifyCommands,
        ConfigureCommands: tt.fields.ConfigureCommands,
      }
      if err := s.Run(tt.args.ex, tt.args.name, tt.args.cmdList); (err != nil) != tt.wantErr {
        t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
      }
    })
  }
}