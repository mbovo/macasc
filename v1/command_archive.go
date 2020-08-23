package v1

import (
	"errors"
	"fmt"

	"github.com/mbovo/yacasc/v1/internal"
)

func Unzip(c *Command) Result {
	if e := internal.VerifyRequiredArgs(c.Name, []string{"src", "dest"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	src := c.args["src"]
	dest := c.args["dest"]

	retVal := Result{Type: CHANGED}

	if internal.FileExists(src.(string)) {
		_, e := internal.Unzip(src.(string), dest.(string))
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
		retVal.Error = errors.New ("src file not found")
	}
	return retVal
}