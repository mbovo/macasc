package v2

import (
  "github.com/mbovo/yacasc/v2/command"
)

type Runnable interface {
  Run(vars command.VarMap, cb command.OutCallback)
}
