package v1

import (
  "reflect"
  "testing"
)

func TestLoadVarFile(t *testing.T) {
  type args struct {
    uri string
  }
  tests := []struct {
    name    string
    args    args
    want    Vars
    wantErr bool
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      got, err := LoadVarFile(tt.args.uri)
      if (err != nil) != tt.wantErr {
        t.Errorf("LoadVarFile() error = %v, wantErr %v", err, tt.wantErr)
        return
      }
      if !reflect.DeepEqual(got, tt.want) {
        t.Errorf("LoadVarFile() got = %v, want %v", got, tt.want)
      }
    })
  }
}