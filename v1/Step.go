package v1

import (
  "gopkg.in/yaml.v2"

  "github.com/mbovo/yacasc/v1/internal"
)

// Step is a single logical named entity
//	Add 				will group all Command to perform in order to change the target system
//	Verify			will group all Command that verify the status of the system in order to choose which changes to made
//	Remove			will group all Command required to rollback the target system
//	Configure 	will group all Command required to configure the target system after Adding changes
type Step struct {
  Name              string
  AddCommands       []map[string]interface{} `yaml:"add"`
  RemoveCommands    []map[string]interface{} `yaml:"remove"`
  VerifyCommands    []map[string]interface{} `yaml:"verify"`
  ConfigureCommands []map[string]interface{} `yaml:"configure"`
}

// Load steps from file (local or remote), context is a map of variables to resolve
func LoadStepFile(uri string, context map[string]interface{}) ([]Step, error) {
  var s []Step
  data, e := internal.LoadFromFileOrUri(uri)
  if e != nil {
    return s, e
  }
  tpl, e := internal.Template(string(data), context)
  if e != nil {
    return s, e
  }
  e = yaml.Unmarshal([]byte(tpl), &s)
  if e != nil {
    return s, e
  }

  return s, nil
}

func (s Step)Run(ex Executor, name string, cmdList []map[string]interface{}) error {

  for _, cmdStr := range cmdList {
    var result Result
    if cmdName, ok := cmdStr["name"].(string); ok {

      found := false
      for _, cmd := range ex.Commands {
        if cmd.Is(cmdName) {
          ex.callback.Output("- %10s  ", cmdName)
          result = cmd.Execute(cmdStr, ex.Vars, ex.callback)

          if varName, ok := cmdStr["result"].(string); ok {
            ex.Vars[varName] = result.Info
          }
          found = true
          break
        }
      }

      if !found {
        ex.callback.Error("command %s not found", cmdName)
      }

    } else {
      ex.callback.Error("cannot decode command name for %#v", cmdStr)
      continue
    }
    ex.callback.Result(result)

  }
  return nil
}
