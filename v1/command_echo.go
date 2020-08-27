package v1

import (
  "fmt"
  "strings"
)

func Echo(c *Command) Result {

  args, r := GetStringArrayArgument(c, "args")
  if r != nil { return *r }

  retVal := Result{Type: OK}
  builder := strings.Builder{}

  for _, arg := range args {
    if arg != "" {
      builder.WriteString(fmt.Sprintf("%s\n", arg))
    }
  }
  retVal.Message = builder.String()
  return retVal

}
