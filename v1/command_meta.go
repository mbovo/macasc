package v1

//func Configure(ex *executor.Executor, step *executor.Step) CommandResult{
//
//	retVal := CommandResult{Type: ResultOK}
//
//	if e:=step.Run(*ex, "Configure", step.ConfigureCommands); e!= nil {
//		retVal.Type = ResultERROR
//		retVal.Error = e
//	}
//	return retVal
//}
//
//func (c Command) Verify(ex *executor.Executor, step *executor.Step) CommandResult {
//
//	retVal := CommandResult{Type: ResultOK}
//
//	if e := step.Run(*ex, "Verify", step.VerifyCommands); e != nil {
//		retVal.Type = ResultERROR
//		retVal.Error = e
//	}
//	return retVal
//}
