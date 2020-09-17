package v2

import (
  "strings"

  "github.com/mbovo/yacasc/v2/command"
)

type cmd struct {
 name string
 run func(varMap command.VarMap) command.Result
 aliases []string
 helpMessage string
}

func (c cmd) Name() string {
 return c.name
}

func (c cmd) Aliases() []string {
 return c.aliases
}

// Return boolean when name given match Name or Aliases
func (c cmd) Is(name string) bool {
 if strings.EqualFold(c.Name(), name) {
   return true
 }

 for _, alas := range c.Aliases() {
   if strings.EqualFold(alas, name) {
     return true
   }
 }
 return false
}

func (c cmd) Exec(vars command.VarMap) command.Result {
 return c.run(vars)
}

func (c cmd) Help(callback command.OutCallback) {
 callback.Output(c.helpMessage)
}
