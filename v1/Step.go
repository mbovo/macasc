package v1

import (
  "strings"

  "gopkg.in/flosch/pongo2.v3"
  "gopkg.in/yaml.v2"

  "github.com/mbovo/yacasc/v1/internal"
)

// Step is a single logical named entity
type Step struct {
  Name     string
  Vars     map[string]interface{}
  Commands []map[string]map[string]interface{} `yaml:"cmds"`
}

// Load steps from file (local or remote), context is a map of variables to resolve
func LoadStepFile(uri string, context map[string]interface{}) ([]Step, error) {
  var s []Step
  data, e := internal.LoadFromFileOrUri(uri)
  if e != nil {
    return s, e
  }

  e = yaml.Unmarshal(data, &s)

  ctx := make(map[string]interface{})
  for k, v := range context{
    ctx[k] = v
  }
  for _, step := range s {
    for k, v := range step.Vars{
      ctx[k] = v
    }
  }

  tpl, e := internal.Template(string(data),ctx )
  if e != nil {
    return s, e
  }
  e = yaml.Unmarshal([]byte(tpl), &s)
  if e != nil {
    return s, e
  }

  return s, nil
}

func (s Step) Run(ex Executor) error {

  // iterate over command in this step
  for _, cmdMap := range s.Commands {

    var result Result
    found := false

    // each command is a map, iterate over its values
    for cmdName, cmdArgs := range cmdMap {

      // iterate over available command, looking for the right one
      for _, cmd := range ex.Commands {

        // verify if name or aliases matches
        if cmd.Is(cmdName) {

          ex.callback.Output("- %10s  ", cmdName)
          // this command must be skipped?
          if r, ok := mustSkip(cmdMap, ex); ok{
            result = r
            goto skip
          }

          // exec this command
          result = cmd.Execute( cmdArgs , ex.Vars, ex.callback)

          // optionally register output var
          registerVar(ex, cmdMap, result)
          found = true
          break
        }

        //Workaround: jump out if the "command" is "opts"
        if strings.EqualFold(cmdName, "opts") {
          found = true
          continue
        }
      }
      if !found {
        ex.callback.Error("command %s not found", cmdName)
      }
    }

  skip:
    ex.callback.Result(result)
  }
  return nil
}

func registerVar(ex Executor, cmdMap map[string]map[string]interface{}, result Result) {
  // if opts.register is defined, save the result in global vars
  if opts, ok := cmdMap["opts"]; ok {
    if varName, ok := opts["register"]; ok {
      ex.Vars[varName.(string)] = result.Info
    }
  }
}


func mustSkip(cmdMap map[string]map[string]interface{}, ex Executor) (result Result, ok bool){

  if opts, ok := cmdMap["opts"]; ok {
    if condition, ok := opts["when"]; ok {

      switch condition.(type) {
      case bool:
        if !condition.(bool) {
          result.Type = SKIPPED
          return result, true
        }
      case string:
        r, err := internal.Template(condition.(string), pongo2.Context(ex.Vars))
        if err != nil {
          ex.callback.Error("%s\n", err)
          return Result{}, false
        }
        if strings.EqualFold(r,"false") || strings.EqualFold(r, ""){
          result.Type = SKIPPED
          return result, true
        }
      }
    }
  }
  return result, false
}