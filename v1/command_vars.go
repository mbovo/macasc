package v1

import (
	"errors"
)

func SetVars(c *Command) Result {

	retVal := Result{Type: OK}

	args, ok := c.args["vars"].(map[interface{}]interface{})

	if !ok {
		retVal.Error = errors.New("invalid or empty args given")
		retVal.Type = ERROR
		return retVal
	}

	for k, v := range args {
		c.vars[k.(string)] = v
	}

	return retVal
}

func PrintVars(c *Command) Result {

	retVal := Result{Type: OK}
	retVal.Info = c.vars

	return retVal
}
