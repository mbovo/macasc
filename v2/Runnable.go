package v2

type Runnable interface {
  Run(vars VarMap, cb OutCallback)
}

type JobRunner struct {
  
}

func (j JobRunner) Run(vars VarMap, cb OutCallback) {
  panic("implement me")
}
