package v2

import (
  "reflect"
  "testing"

  "github.com/mbovo/yacasc/v2/command"
)

func TestNewResult(t *testing.T) {
  type args struct {
    status  string
    message string
    info    map[string]string
  }

  tests := []struct {
    name string
    args args
    want command.Result
  }{
    {name: "Empty one", args: args{}, want: result{}},
    {name: "Only Message", args: args{OK, "Hallo world", nil}, want: result{OK, "Hallo world", nil}},
    {name: "Only Info", args: args{OK, "", map[string]string{"key": "value"}}, want: result{OK, "", map[string]string{"key": "value"}}},
    {name: "Error", args: args{ERROR, "This is error", nil}, want: result{ERROR, "This is error", nil}},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := NewResult(tt.args.status, tt.args.message, tt.args.info); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("NewResult() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestResult_Error(t *testing.T) {
  type fields struct {
    status  string
    message string
    info    map[string]string
  }
  tests := []struct {
    name   string
    fields fields
    want   string
  }{
    {"Simple error", fields{ERROR, "simple error", nil}, "simple error"},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := result{
        status:  tt.fields.status,
        message: tt.fields.message,
        info:    tt.fields.info,
      }
      if got := c.Error(); got != tt.want {
        t.Errorf("Error() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestResult_HasChanged(t *testing.T) {
  type fields struct {
    status  string
    message string
    info    map[string]string
  }
  tests := []struct {
    name   string
    fields fields
    want   bool
  }{
    {"Changed", fields{CHANGED, "", nil}, true},
    {"OK", fields{OK, "", nil}, false},
    {"Error", fields{ERROR, "", nil}, false},
    {"Skipped", fields{SKIPPED, "", nil}, false},
    {"Empty", fields{}, false},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := result{
        status:  tt.fields.status,
        message: tt.fields.message,
        info:    tt.fields.info,
      }
      if got := c.HasChanged(); got != tt.want {
        t.Errorf("HasChanged() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestResult_HasSkipped(t *testing.T) {
  type fields struct {
    status  string
    message string
    info    map[string]string
  }
  tests := []struct {
    name   string
    fields fields
    want   bool
  }{
    {"Changed", fields{CHANGED, "", nil}, false},
    {"OK", fields{OK, "", nil}, false},
    {"Error", fields{ERROR, "", nil}, false},
    {"Skipped", fields{SKIPPED, "", nil}, true},
    {"Empty", fields{}, false},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := result{
        status:  tt.fields.status,
        message: tt.fields.message,
        info:    tt.fields.info,
      }
      if got := c.HasSkipped(); got != tt.want {
        t.Errorf("HasSkipped() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestResult_Info(t *testing.T) {
  type fields struct {
    status  string
    message string
    info    map[string]string
  }
  tests := []struct {
    name   string
    fields fields
    want   map[string]string
  }{
    {"Changed", fields{CHANGED, "message", map[string]string{"elementkey": "elementval"}}, map[string]string{"elementkey": "elementval"}},
    {"OK", fields{OK, "message", map[string]string{"elementkey": "elementval"}}, map[string]string{"elementkey": "elementval"}},
    {"Error", fields{ERROR, "message", map[string]string{"elementkey": "elementval"}}, map[string]string{"elementkey": "elementval"}},
    {"Skipped", fields{SKIPPED, "message", map[string]string{"elementkey": "elementval"}}, map[string]string{"elementkey": "elementval"}},
    {"Empty", fields{}, nil},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := result{
        status:  tt.fields.status,
        message: tt.fields.message,
        info:    tt.fields.info,
      }
      if got := c.Info(); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Info() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestResult_IsErr(t *testing.T) {
  type fields struct {
    status  string
    message string
    info    map[string]string
  }
  tests := []struct {
    name   string
    fields fields
    want   bool
  }{
    {"Changed", fields{CHANGED, "", nil}, false},
    {"OK", fields{OK, "", nil}, false},
    {"Error", fields{ERROR, "", nil}, true},
    {"Skipped", fields{SKIPPED, "", nil}, false},
    {"Empty", fields{}, false},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := result{
        status:  tt.fields.status,
        message: tt.fields.message,
        info:    tt.fields.info,
      }
      if got := c.IsErr(); got != tt.want {
        t.Errorf("IsErr() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestResult_IsOk(t *testing.T) {
  type fields struct {
    status  string
    message string
    info    map[string]string
  }
  tests := []struct {
    name   string
    fields fields
    want   bool
  }{
    {"Changed", fields{CHANGED, "", nil}, false},
    {"OK", fields{OK, "", nil}, true},
    {"Error", fields{ERROR, "", nil}, false},
    {"Skipped", fields{SKIPPED, "", nil}, false},
    {"Empty", fields{}, false},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := result{
        status:  tt.fields.status,
        message: tt.fields.message,
        info:    tt.fields.info,
      }
      if got := c.IsOk(); got != tt.want {
        t.Errorf("IsOk() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestResult_IsZero(t *testing.T) {
  type fields struct {
    status  string
    message string
    info    map[string]string
  }
  tests := []struct {
    name   string
    fields fields
    want   bool
  }{
    {"Changed", fields{CHANGED, "", nil}, false},
    {"OK", fields{OK, "", nil}, false},
    {"Error", fields{ERROR, "", nil}, false},
    {"Skipped", fields{SKIPPED, "", nil}, false},
    {"Empty", fields{}, true},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := result{
        status:  tt.fields.status,
        message: tt.fields.message,
        info:    tt.fields.info,
      }
      if got := c.IsZero(); got != tt.want {
        t.Errorf("IsZero() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestResult_Message(t *testing.T) {
  type fields struct {
    status  string
    message string
    info    map[string]string
  }
  tests := []struct {
    name   string
    fields fields
    want   string
  }{
    {"Changed", fields{CHANGED, "message", nil}, "message"},
    {"OK", fields{OK, "message", nil}, "message"},
    {"Error", fields{ERROR, "message", nil}, "message"},
    {"Skipped", fields{SKIPPED, "message", nil}, "message"},
    {"Empty", fields{}, ""},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := result{
        status:  tt.fields.status,
        message: tt.fields.message,
        info:    tt.fields.info,
      }
      if got := c.Message(); got != tt.want {
        t.Errorf("Message() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestResult_Status(t *testing.T) {
  type fields struct {
    status  string
    message string
    info    map[string]string
  }
  tests := []struct {
    name   string
    fields fields
    want   string
  }{
    {"Changed", fields{CHANGED, "message", nil}, CHANGED},
    {"OK", fields{OK, "message", nil}, OK},
    {"Error", fields{ERROR, "message", nil}, ERROR},
    {"Skipped", fields{SKIPPED, "message", nil}, SKIPPED},
    {"Empty", fields{}, ""},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := result{
        status:  tt.fields.status,
        message: tt.fields.message,
        info:    tt.fields.info,
      }
      if got := c.Status(); got != tt.want {
        t.Errorf("Status() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestResult_String(t *testing.T) {
  type fields struct {
    status  string
    message string
    info    map[string]string
  }
  tests := []struct {
    name   string
    fields fields
    want   string
  }{
    {"Changed", fields{CHANGED, "message", map[string]string{"key": "value"}}, "[CHANGED]: message\nkey: value"},
    {"OK", fields{OK, "message", map[string]string{"key": "value"}}, "[OK]: message\nkey: value"},
    {"Error", fields{ERROR, "message", map[string]string{"key": "value"}}, "[ERROR]: message\nkey: value"},
    {"Skipped", fields{SKIPPED, "message", map[string]string{"key": "value"}}, "[SKIPPED]: message\nkey: value"},
    {"Empty", fields{}, ""},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      c := result{
        status:  tt.fields.status,
        message: tt.fields.message,
        info:    tt.fields.info,
      }
      if got := c.String(); got != tt.want {
        t.Errorf("String() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestReturnChanged(t *testing.T) {
  type args struct {
    message string
    info    map[string]string
  }
  tests := []struct {
    name string
    args args
    want command.Result
  }{
    {"With message", args{"message", nil}, result{CHANGED, "message", nil}},
    {"With Info", args{"", map[string]string{"key": "value"}}, result{CHANGED, "", map[string]string{"key": "value"}}},
    {"With Both", args{"message", map[string]string{"key": "value"}}, result{CHANGED, "message", map[string]string{"key": "value"}}},
    {"Empty", args{}, result{CHANGED, "", nil}},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := ReturnChanged(tt.args.message, tt.args.info); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("ReturnChanged() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestReturnError(t *testing.T) {
  type args struct {
    message string
    info    map[string]string
  }
  tests := []struct {
    name string
    args args
    want command.Result
  }{
    {"With message", args{"message", nil}, result{ERROR, "message", nil}},
    {"With Info", args{"", map[string]string{"key": "value"}}, result{ERROR, "", map[string]string{"key": "value"}}},
    {"With Both", args{"message", map[string]string{"key": "value"}}, result{ERROR, "message", map[string]string{"key": "value"}}},
    {"Empty", args{}, result{ERROR,"", nil}},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := ReturnError(tt.args.message, tt.args.info); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("ReturnError() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestReturnOk(t *testing.T) {
  type args struct {
    message string
    info    map[string]string
  }
  tests := []struct {
    name string
    args args
    want command.Result
  }{
    {"With message", args{"message", nil}, result{OK, "message", nil}},
    {"With Info", args{"", map[string]string{"key": "value"}}, result{OK, "", map[string]string{"key": "value"}}},
    {"With Both", args{"message", map[string]string{"key": "value"}}, result{OK, "message", map[string]string{"key": "value"}}},
    {"Empty", args{}, result{OK,"", nil}},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := ReturnOk(tt.args.message, tt.args.info); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("ReturnOk() = %v, want %v", got, tt.want)
      }
    })
  }
}

func TestReturnSkipped(t *testing.T) {
  type args struct {
    message string
    info    map[string]string
  }
  tests := []struct {
    name string
    args args
    want command.Result
  }{
    {"With message", args{"message", nil}, result{SKIPPED, "message", nil}},
    {"With Info", args{"", map[string]string{"key": "value"}}, result{SKIPPED, "", map[string]string{"key": "value"}}},
    {"With Both", args{"message", map[string]string{"key": "value"}}, result{SKIPPED, "message", map[string]string{"key": "value"}}},
    {"Empty", args{}, result{SKIPPED,"", nil}},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := ReturnSkipped(tt.args.message, tt.args.info); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("ReturnSkipped() = %v, want %v", got, tt.want)
      }
    })
  }
}
