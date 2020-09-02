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
  callback OutputCallback
  Steps    []Step
  Vars     Vars
  Commands []Command
}

type ExecutorBuilder struct {
  cb       OutputCallback
  varFile  string
  stepFile string
  cmds     []Command
}

func (eb *ExecutorBuilder) AddStepsFromFile( filename string) *ExecutorBuilder{
  eb.stepFile = filename
  return eb
}

func (eb *ExecutorBuilder) AddVarsFromFile( filename string ) *ExecutorBuilder{
  eb.varFile = filename
  return eb
}

func (eb *ExecutorBuilder) AddCommands( cmds []Command) *ExecutorBuilder {
  eb.cmds = cmds
  return eb
}

func (eb *ExecutorBuilder) AddCallback( cb OutputCallback) *ExecutorBuilder{
  eb.cb = cb
  return eb
}

func (eb *ExecutorBuilder) Build() (*Executor, error) {

  if eb.stepFile == "" {
    return nil, fmt.Errorf("invalid data, ExecutorBuilder needs a step file")
  }

  v, e := LoadVarFile(eb.varFile)
  if e != nil {
  return nil, e
  }

  s, e := LoadStepFile(eb.stepFile, v)
  if e != nil {
  return nil, e
  }

  return NewExecutor(eb.cb, v, s, eb.cmds)
}


// Returns executor object loading Steps and Vars from given filenames
// It accept a callback function in order to write back status to UX
func NewExecutor(callback OutputCallback, v Vars, s []Step, cmds []Command) (*Executor, error) {

  // add vars declared on each step to global vars
  for _, steps := range s {
    for key, value := range steps.Vars {
      v[key] = value
    }
  }

  ex := Executor{
    callback: callback,
    Vars:     v,
    Steps:    s,
    Commands: cmds,
  }

  return &ex, nil
}

// Run each step attached to this executor
func (ex Executor) Run() (err error) {

  ex.callback.Output("Loaded %s \n", ex.StepsFile)

  tot := len(ex.Steps)

  for i, step := range ex.Steps {
    ex.callback.Output("Step %d/%d: %s\n", i+1, tot, step.Name)
    err = step.Run(ex)
    if err != nil {
      return
    }
  }

  return err
}
