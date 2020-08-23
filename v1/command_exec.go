package v1

import (
	"github.com/mbovo/yacasc/v1/internal"
)

func Exec(c *Command) Result {

	if e := internal.VerifyRequiredArgs(c.Name, []string{"cmd", "args"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	cmd := c.args["cmd"]
	args := c.args["args"].([]string)

	retVal := Result{Type: CHANGED}

	//TODO: return stderr and stdout content and write it to retVal.Info
	if _, err := internal.Exec(false, cmd.(string), args...); err != nil {
		retVal.Error = err
		retVal.Type = ERROR
	}

	return retVal
}

func Shell(c *Command) Result {

	if e := internal.VerifyRequiredArgs(c.Name, []string{"args"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	args := c.args["args"].([]string)
	cd, ok := c.args["cwd"]
	if !ok {
		cd = ""
	}

	retVal := Result{Type: CHANGED}

	//TODO: return stderr and stdout content and write it to retVal.Info
	if _, err := internal.Shell(cd.(string), args...); err != nil {
		retVal.Error = err
		retVal.Type = ERROR
	}

	return retVal
}
