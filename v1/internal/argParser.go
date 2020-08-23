package internal

import (
	"fmt"
)

// Given a command, its name and a list of required keys
// returns an error if any of the keys are not in the command's arguments
func VerifyRequiredArgs(cmdName string , requiredKeys []string, args map[string]interface{}) error {
	for _, k := range requiredKeys{
		e, ok := args[k]
		if !ok || e == nil {
			return fmt.Errorf("command %s: missing argument: %s\n", cmdName, k)
		}
	}
	return nil
}