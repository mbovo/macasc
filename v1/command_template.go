package v1

import (
	"errors"
	"fmt"

	"github.com/mbovo/yacasc/v1/internal"
)

func Template(c *Command) Result {

	src, r := GetStringArgument(c, "src")
	if r != nil {
		return *r
	}
	dest, r := GetStringArgument(c, "dest")
	if r != nil {
		return *r
	}

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
	retVal.Message = fmt.Sprintf("%s -> %s", src, dest)
	return retVal
}
