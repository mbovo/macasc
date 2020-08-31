package v1

import (
  "gopkg.in/yaml.v2"

  "github.com/mbovo/yacasc/v1/internal"
)

// Step is a single logical named entity
type Step struct {
  Name     string
  Commands []map[string]map[string]interface{} `yaml:"cmds"`
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

func (s Step) Run(ex Executor, name string, cmdList []map[string]interface{}) error {

  // iterate over command in this step
  for _, cmdMap := range s.Commands {

    var result Result
    found := false

    // each command is a map, iterate over its values
    for cmdName, cmdArgs := range cmdMap {

      // iterate over available commands, looking for the right one
      for _, cmd := range ex.Commands {

        // verify if name or aliases matches
        if cmd.Is(cmdName) {

          ex.callback.Output("- %10s  ", cmdName)
          // exec this command
          result = cmd.Execute( cmdArgs , ex.Vars, ex.callback)

          //todo: re-implement register variables

          found = true
          break
        }
      }
      if !found {
        ex.callback.Error("command %s not found", cmdName)
      }
    }
    ex.callback.Result(result)
  }
  return nil
}
