package v1

import (
  "reflect"
  "testing"
)

func TestCopy(t *testing.T) {
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
      if got := Copy(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Copy() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestFileExists(t *testing.T) {
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
      if got := FileExists(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("FileExists() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestFind(t *testing.T) {
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
      if got := Find(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Find() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestLink(t *testing.T) {
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
      if got := Link(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Link() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestMkdir(t *testing.T) {
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
      if got := Mkdir(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Mkdir() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestMove(t *testing.T) {
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
      if got := Move(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Move() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestRemoveFiles(t *testing.T) {
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
      if got := RemoveFiles(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("RemoveFiles() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestWhich(t *testing.T) {
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
      if got := Which(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Which() = %v, want %v", got, tt.want)
      }
    })
  }
}
