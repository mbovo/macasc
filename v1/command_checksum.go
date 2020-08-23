package v1

import (
	"github.com/mbovo/yacasc/v1/internal"
)

func Hashfile(c *Command) Result {

	if e := internal.VerifyRequiredArgs(c.Name, []string{"args"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	args := c.args["args"]

	retVal := Result{Type: OK}

	info := make(map[string]interface{})

	for _, path := range args.([]interface{}) {
		if path == nil { continue }
		h, e := internal.SHA256File(path.(string))
		if e != nil {
			retVal.Error = e
			retVal.Type = ERROR
			info[path.(string)] = e.Error()
		}
		info[path.(string)] = h
	}
	retVal.Info = info
	return retVal

}

func HashString(c *Command) Result {

	if e := internal.VerifyRequiredArgs(c.Name, []string{"args"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	args := c.args["args"]

	retVal := Result{Type: OK}

	info := make(map[string]interface{})

	for _, str := range args.([]interface{}) {
		h := internal.SHA256String(str.(string))
		info[str.(string)] = h
	}
	retVal.Info = info
	return retVal
}