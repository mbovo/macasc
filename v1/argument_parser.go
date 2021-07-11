package v1

import (
	"fmt"
)

//Given a command and a list of required keys
//returns an Result of type ERROR if any of the keys requested is missing
func VerifyRequiredArgs(c *Command, required []string) *Result {
	for _, arg := range required {
		_, r := getArgument(c, arg)
		if r != nil {
			return r
		}
	}
	return nil
}

func getArgument(c *Command, name string) (interface{}, *Result) {
	i, ok := c.args[name]
	if !ok {
		return i, &Result{Type: ERROR, Error: fmt.Errorf("required argument not found. %s", name)}
	}
	return i, nil
}

func GetStringArrayArgument(c *Command, argumentName string) (l []string, r *Result) {
	v, r := getArgument(c, argumentName)
	if r != nil {
		return
	}

	switch t := v.(type) {
	case []interface{}:
		list := v.([]interface{})
		for _, value := range list {
			l = append(l, value.(string))
		}
	case []string:
		for _, value := range v.([]string) {
			l = append(l, value)
		}
	default:
		return l, &Result{Type: ERROR, Error: fmt.Errorf("argument %s is of invalid type, must be []string, found %T", argumentName, t)}
	}

	return l, nil
}

func GetStringArgument(c *Command, argumentName string) (s string, r *Result) {
	v, r := getArgument(c, argumentName)
	if r != nil {
		return
	}

	switch t := v.(type) {
	case string:
		s = v.(string)
	default:
		return s, &Result{Type: ERROR, Error: fmt.Errorf("argument %s is of invalid type, must be string, found %T", argumentName, t)}
	}
	return s, nil
}

func GetIntArgument(c *Command, argumentName string) (n int, r *Result) {
	v, r := getArgument(c, argumentName)
	if r != nil {
		return
	}

	switch t := v.(type) {
	case int:
		n = v.(int)
	default:
		return n, &Result{Type: ERROR, Error: fmt.Errorf("argument %s is of invalid type, must be string, found %T", argumentName, t)}
	}
	return n, nil
}

func GetBoolArgument(c *Command, argumentName string) (b bool, r *Result) {
	v, r := getArgument(c, argumentName)
	if r != nil {
		return
	}

	switch t := v.(type) {
	case int:
		b = v.(bool)
	default:
		return b, &Result{Type: ERROR, Error: fmt.Errorf("argument %s is of invalid type, must be string, found %T", argumentName, t)}
	}
	return b, nil
}

func GetArgumentFromVars(c *Command) (slice []string, r *Result) {
	v, r := getArgument(c, "fromVar")
	if r != nil {
		return slice, r
	}
	varName := v.(string)
	varValue := c.vars[varName]

	switch varValue.(type) {
	case string:
		slice = append(slice, varValue.(string))
	case []interface{}:
		slice = append(slice, varValue.([]string)...)
	case map[string]interface{}:
		for k, _ := range varValue.(map[string]interface{}) {
			slice = append(slice, k)
		}
	}
	return slice, nil
}
