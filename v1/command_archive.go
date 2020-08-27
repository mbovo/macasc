package v1

import (
  "errors"
  "fmt"

  "github.com/mbovo/yacasc/v1/internal"
)

func Unzip(c *Command) Result {

  src, res := GetStringArgument(c, "src")
  if res != nil {
    return *res
  }
  dest, res := GetStringArgument(c, "src")
  if res != nil {
    return *res
  }

  retVal := Result{Type: CHANGED}

  if internal.FileExists(src) {
    _, e := internal.Unzip(src, dest)
    if e != nil {
      retVal.Type = ERROR
      retVal.Error = e
    }

    retVal.Message = fmt.Sprintf("%s -> %s", src, dest)

    //retVal.Info = map[string]interface{}{}
    //for _, f := range files{
    //	retVal.Info[f] = OK
    //}
  } else {
    retVal.Type = ERROR
    retVal.Error = errors.New("src file not found")
  }
  return retVal
}
