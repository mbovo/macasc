package v1

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

  // add vars declared on each step to global vars
  for _, steps := range s {
    for key, value := range steps.Vars {
      v[key] = value
    }
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
