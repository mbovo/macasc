package v2

type Command interface {
  Exec(vars VarMap) CommandResult
  Help(callback OutCallback)
}

type command struct {
  
}

func (c command) Exec(vars VarMap) CommandResult {
  panic("implement me")
}

func (c command) Help(callback OutCallback) {
  panic("implement me")
}
