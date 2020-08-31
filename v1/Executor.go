package v1

import (
  "fmt"
)

// Keep track of current execution
//	VarFile is the vars filename (can be a remote url)
//	StepsFile is the steps filename ( can be a remote url)
//	Steps	contains currently loaded steps
//	Vars contains currently loaded vars
//	Commands must be a list of valid command.Command
type Executor struct {
  callback  OutputCallback
  VarFile   string
  StepsFile string
  Steps     []Step
  Vars      Vars
  Commands  []Command
}

// An action type
type Action uint8

// Different actions
const (
  Verify Action = 1 * iota
  Apply
  Configure
  Remove
)

// Returns executor object loading Steps and Vars from given filenames
// It accept a callback function in order to write back status to UX
func NewExecutor(callback OutputCallback, StepsFile, VarFile string) (*Executor, error) {

  v, e := LoadVarFile(VarFile)
  if e != nil {
    return nil, e
  }

  s, e := LoadStepFile(StepsFile, v)
  if e != nil {
    return nil, e
  }

  ex := Executor{
    callback:  callback,
    VarFile:   VarFile,
    StepsFile: StepsFile,
    Vars:      v,
    Steps:     s,
    Commands:  DefaultList,
  }

  return &ex, nil
}

// Run each step attached to this executor
func (ex Executor) Run(action Action) (err error) {

  //ex.callback.Output("Loaded %s \n", ex.StepsFile)

  tot := len(ex.Steps)

  for i, step := range ex.Steps {

    ex.callback.Output("Step %d/%d: %s\n", i+1, tot, step.Name)

    switch action {
    case Verify:
      err = step.Run(ex, "Verify", step.VerifyCommands)
      if err != nil {
        return
      }
      break
    case Apply:
      err = step.Run(ex, "Add", step.AddCommands)
      if err != nil {
        return
      }
      break
    case Configure:
      err = step.Run(ex, "Configure", step.ConfigureCommands)
      if err != nil {
        return
      }
      break
    case Remove:
      err = step.Run(ex, "Remove", step.RemoveCommands)
      if err != nil {
        return
      }
      break
    default:
      return fmt.Errorf("invalid action type")

    }
  }
  return err
}
