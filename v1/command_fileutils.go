package v1

import (
	"errors"
	"fmt"
	"os"

	"github.com/mbovo/yacasc/v1/internal"
)


func Link(c *Command) Result {
	if e := internal.VerifyRequiredArgs(c.Name, []string{"src", "dest"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	src := c.args["src"]
	dest := c.args["dest"]

	retVal := Result{}

	if !internal.FileExists(src.(string)) {
		retVal.Type = ERROR
		retVal.Error = errors.New("src file does not exists")
		return retVal
	}
	if internal.FileExists(dest.(string)) {
		retVal.Type = ERROR
		retVal.Error = errors.New("dest file already exists")
		return retVal
	}

	if e := os.Link(src.(string), dest.(string)); e != nil {
		retVal.Type = ERROR
		retVal.Error = e

	} else {
		retVal.Type = CHANGED
		retVal.Message = fmt.Sprintf("Linked %s to %s", src.(string), dest.(string))
	}
	return retVal
}

func RemoveFiles(c *Command) Result {

	if e := internal.VerifyRequiredArgs(c.Name, []string{"args"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	fileList := c.args["args"]

	retVal := Result{Type: CHANGED}

	info := make(map[string]interface{})

	for _, path := range fileList.([]interface{}) {
		if path == nil { continue }
		if err := os.RemoveAll(path.(string)); err != nil {
			retVal.Error = err
			retVal.Type = ERROR
			info[path.(string)] = err.Error()
		}
		info[path.(string)] = ""
	}
	retVal.Info = info
	return retVal
}



func Mkdir(c *Command) Result {
	if e := internal.VerifyRequiredArgs(c.Name, []string{"args"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	args := c.args["args"]

	mode, ok := c.args["mode"]
	if !ok {
		mode = os.ModePerm
	}

	retVal := Result{Type: CHANGED}

	info := make(map[string]interface{})

	for _, path := range args.([]interface{}) {
		fi, e := os.Stat(path.(string))
		if e == nil && fi.IsDir() {
			info[path.(string)] = OK
		}
		if e := internal.MakeDir(path.(string), mode.(os.FileMode)); e != nil {
			retVal.Type = ERROR
			info[path.(string)] = e.Error()
		}
	}
	retVal.Info = info
	return retVal
}

func FileExists(c *Command) Result {

	if e := internal.VerifyRequiredArgs(c.Name, []string{"args"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	fileList := c.args["args"]
	retVal := Result{Type: OK}

	info := make(map[string]interface{})

	for _, path := range fileList.([]interface{}) {
		if !internal.FileExists(path.(string)) {
			retVal.Error = errors.New("file not found")
			retVal.Type = ERROR
			info[path.(string)] = "NOT FOUND"
		}
		info[path.(string)] = "FOUND"
	}
	retVal.Info = info
	return retVal
}


func Move(c *Command) Result {
	if e := internal.VerifyRequiredArgs(c.Name, []string{"src", "dest"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	src := c.args["src"]
	dest := c.args["dest"]
	force, ok := c.args["force"]
	if !ok {
		force = false
	}

	retVal := Result{Type: CHANGED}

	if internal.FileExists(dest.(string)) {
		// destination path exists
		if !(force.(bool)) {
			// and force is false, return an error
			retVal.Type = ERROR
			retVal.Error = errors.New("destination file already exists")
			return retVal
		}
	}
	if !internal.FileExists(src.(string)) {
		// source path doesn't exists, exit
		retVal.Type = ERROR
		retVal.Error = errors.New("source file does not exists")
		return retVal
	}

	// rename the file
	if e := os.Rename(src.(string), dest.(string)); e != nil {
		retVal.Type = ERROR
		retVal.Error = e
	}

	retVal.Message = fmt.Sprintf("%s -> %s", src, dest)

	return retVal
}

func Copy(c *Command) Result {
	if e := internal.VerifyRequiredArgs(c.Name, []string{"src", "dest"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	src := c.args["src"]
	dest := c.args["dest"]
	force, ok := c.args["force"]
	if !ok {
		force = false
	}

	retVal := Result{Type: CHANGED}

	if !internal.FileExists(src.(string)) {
		// source path doesn't exists, exit
		retVal.Type = ERROR
		retVal.Error = errors.New("source file does not exists")
		return retVal
	}
	if internal.FileExists(dest.(string)) {
		// destination path exists
		if !(force.(bool)) {
			// and force is false, return an error
			retVal.Type = ERROR
			retVal.Error = errors.New("destination file already exists (and 'force' was false)")
			return retVal
		}
		// force=true, before copy, calculate both hashes
		srcH, _ := internal.SHA256File (src.(string))
		dstH, _ := internal.SHA256File(dest.(string))
		if srcH == dstH {
			// nothing to copy
			retVal.Type = OK
			retVal.Message = fmt.Sprintf("%s == %s", src, dest)
			return retVal
		}
	}

	// copy the file
	if e := internal.CopyFile(src.(string), dest.(string)); e != nil {
		retVal.Type = ERROR
		retVal.Error = e
	}

	retVal.Message = fmt.Sprintf("%s -> %s", src, dest)

	return retVal
}



func Which(c *Command) Result {

	if e := internal.VerifyRequiredArgs(c.Name, []string{"args"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	args := c.args["args"].([]interface{})

	retVal := Result{Type: OK, Info: make(map[string]interface{},0)}


	//TODO: return stderr and stdout content and write it to retVal.Info
	for _, name := range args {
		path, ok := internal.ExistsInPath(name.(string))
		if !ok {
			retVal.Error = errors.New("not found in PATH")
			retVal.Type = ERROR
			retVal.Info[name.(string)] = "NOT FOUND IN PATH"
		}
		retVal.Info[name.(string)] = path
	}

	return retVal
}