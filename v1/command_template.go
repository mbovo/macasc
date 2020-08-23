package v1

import (
	"errors"

	"github.com/mbovo/yacasc/v1/internal"
)

func Template(c *Command) Result {

	if e := internal.VerifyRequiredArgs(c.Name, []string{"src", "dest"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	src := c.args["src"].(string)
	dest := c.args["dest"].(string)

	retVal := Result{Type: CHANGED}

	if !internal.FileExists(src) {
		// source path doesn't exists, exit
		retVal.Type = ERROR
		retVal.Error = errors.New("template (src) file does not exists")
		return retVal
	}

	var context map[string]interface{}
	context = c.vars

	if err := internal.TemplateFile(src, dest, context); err != nil {
		retVal.Type = ERROR
		retVal.Error = err
	}
	return retVal
}
