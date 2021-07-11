package v1

import (
	"github.com/mbovo/yacasc/v1/internal"
)

func Hashfile(c *Command) (res Result) {

	// Load command arguments
	args, result := GetStringArrayArgument(c, "args")
	if result != nil {
		return *result
	}

	vars, _ := GetArgumentFromVars(c)

	args = append(args, vars...)

	// Prepare returned values
	res.Type = OK
	info := make(map[string]interface{})

	// Calculate SHA256 for each filename
	for _, path := range args {
		if path == "" {
			continue
		}
		h, e := internal.SHA256File(path)
		if e != nil {
			res.Error = e
			res.Type = ERROR
			info[path] = e.Error()
		}
		info[path] = h
	}
	res.Info = info
	return
}

func HashString(c *Command) Result {

	args, res := GetStringArrayArgument(c, "args")
	if res != nil {
		return *res
	}

	retVal := Result{Type: OK}

	info := make(map[string]interface{})

	for _, str := range args {
		h := internal.SHA256String(str)
		info[str] = h
	}
	retVal.Info = info
	return retVal
}
