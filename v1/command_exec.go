package v1

import (
  "github.com/mbovo/yacasc/v1/internal"
)

func Exec(c *Command) Result {

  cmd, r := GetStringArgument(c, "cmd")
  if r != nil { return *r }

  args, r := GetStringArrayArgument(c, "args")
  if r != nil{ return *r}

  silent, _ := GetBoolArgument(c, "silent")

  retVal := Result{Type: CHANGED}

  //TODO: return stderr and stdout content and write it to retVal.Info
  if _, err := internal.Exec(silent, cmd, args...); err != nil {
    retVal.Error = err
    retVal.Type = ERROR
  }

  return retVal
}

func Shell(c *Command) Result {

  args,r  := GetStringArrayArgument(c, "args")
  if r != nil { return *r}

  cd, _ := GetStringArgument(c, "cwd")

  retVal := Result{Type: CHANGED}

  //TODO: return stderr and stdout content and write it to retVal.Info
  if _, err := internal.Shell(cd, args...); err != nil {
    retVal.Error = err
    retVal.Type = ERROR
  }

  return retVal
}
