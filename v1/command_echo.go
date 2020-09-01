package v1

import (
  "fmt"
  "strings"
)

func Echo(c *Command) Result {

  msg, _ := GetStringArgument(c, "msg")
  args, _ := GetStringArrayArgument(c, "args")

  slice, _ := GetArgumentFromVars(c)
  args = append(args, slice...)

  retVal := Result{Type: OK}
  builder := strings.Builder{}
  builder.WriteString(fmt.Sprintf("%s\n", msg))

  for _, arg := range args {
    if arg != "" {
      builder.WriteString(fmt.Sprintf("%s\n", arg))
    }
  }
  retVal.Message = builder.String()
  return retVal

}
