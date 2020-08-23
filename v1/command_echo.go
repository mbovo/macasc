package v1

import (
	"fmt"
	"strings"

	"github.com/mbovo/yacasc/v1/internal"
)

func Echo(c *Command) Result {

	if e := internal.VerifyRequiredArgs(c.Name, []string{"args"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	args := c.args["args"]

	retVal := Result{Type: OK}
	builder := strings.Builder{}

	for _, arg := range args.([]interface{}) {
		if arg != nil {
			builder.WriteString(fmt.Sprintf("%s\n", arg.(string)))
		}
	}
	retVal.Message = builder.String()
	return retVal

}
