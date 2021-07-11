package v1

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/mbovo/yacasc/v1/internal"
)

func Link(c *Command) Result {

	src, r := GetStringArgument(c, "src")
	if r != nil {
		return *r
	}
	dest, r := GetStringArgument(c, "dest")
	if r != nil {
		return *r
	}

	retVal := Result{}

	if !internal.FileExists(src) {
		retVal.Type = ERROR
		retVal.Error = errors.New("src file does not exists")
		return retVal
	}
	if internal.FileExists(dest) {
		retVal.Type = ERROR
		retVal.Error = errors.New("dest file already exists")
		return retVal
	}

	if e := os.Link(src, dest); e != nil {
		retVal.Type = ERROR
		retVal.Error = e

	} else {
		retVal.Type = CHANGED
		retVal.Message = fmt.Sprintf("Linked %s to %s", src, dest)
	}
	return retVal
}

func RemoveFiles(c *Command) Result {

	fileList, r := GetStringArrayArgument(c, "args")
	if r != nil {
		return *r
	}

	retVal := Result{Type: CHANGED}

	info := make(map[string]interface{})

	for _, path := range fileList {
		if path == "" {
			continue
		}
		if err := os.RemoveAll(path); err != nil {
			retVal.Error = err
			retVal.Type = ERROR
			info[path] = err.Error()
		}
		info[path] = ""
	}
	retVal.Info = info
	return retVal
}

func Mkdir(c *Command) Result {

	args, r := GetStringArrayArgument(c, "args")
	if r != nil {
		return *r
	}

	mode, ok := c.args["mode"]
	if !ok {
		mode = os.ModePerm
	}

	retVal := Result{Type: CHANGED}

	info := make(map[string]interface{})

	for _, path := range args {
		fi, e := os.Stat(path)
		if e == nil && fi.IsDir() {
			info[path] = OK
		}
		if e := internal.MakeDir(path, mode.(os.FileMode)); e != nil {
			retVal.Type = ERROR
			info[path] = e.Error()
		}
	}
	retVal.Info = info
	return retVal
}

func FileExists(c *Command) Result {

	fileList, r := GetStringArrayArgument(c, "args")
	if r != nil {
		return *r
	}

	retVal := Result{Type: OK}

	info := make(map[string]interface{})

	for _, path := range fileList {
		if !internal.FileExists(path) {
			retVal.Error = errors.New("file not found")
			retVal.Type = ERROR
			info[path] = "NOT FOUND"
		}
		info[path] = "FOUND"
	}
	retVal.Info = info
	return retVal
}

func Move(c *Command) Result {

	src, r := GetStringArgument(c, "src")
	if r != nil {
		return *r
	}
	dest, r := GetStringArgument(c, "dest")
	if r != nil {
		return *r
	}

	force, _ := GetBoolArgument(c, "force")

	retVal := Result{Type: CHANGED}

	if internal.FileExists(dest) {
		// destination path exists
		if !(force) {
			// and force is false, return an error
			retVal.Type = ERROR
			retVal.Error = errors.New("destination file already exists")
			return retVal
		}
	}
	if !internal.FileExists(src) {
		// source path doesn't exists, exit
		retVal.Type = ERROR
		retVal.Error = errors.New("source file does not exists")
		return retVal
	}

	// rename the file
	if e := os.Rename(src, dest); e != nil {
		retVal.Type = ERROR
		retVal.Error = e
	}

	retVal.Message = fmt.Sprintf("%s -> %s", src, dest)

	return retVal
}

func Copy(c *Command) Result {

	src, r := GetStringArgument(c, "src")
	if r != nil {
		return *r
	}
	dest, r := GetStringArgument(c, "dest")
	if r != nil {
		return *r
	}

	force, _ := GetBoolArgument(c, "force")

	retVal := Result{Type: CHANGED}

	if !internal.FileExists(src) {
		// source path doesn't exists, exit
		retVal.Type = ERROR
		retVal.Error = errors.New("source file does not exists")
		return retVal
	}
	if internal.FileExists(dest) {
		// destination path exists
		if !(force) {
			// and force is false, return an error
			retVal.Type = ERROR
			retVal.Error = errors.New("destination file already exists (and 'force' was false)")
			return retVal
		}
		// force=true, before copy, calculate both hashes
		srcH, _ := internal.SHA256File(src)
		dstH, _ := internal.SHA256File(dest)
		if srcH == dstH {
			// nothing to copy
			retVal.Type = OK
			retVal.Message = fmt.Sprintf("%s == %s", src, dest)
			return retVal
		}
	}

	// copy the file
	if e := internal.CopyFile(src, dest); e != nil {
		retVal.Type = ERROR
		retVal.Error = e
	}

	retVal.Message = fmt.Sprintf("%s -> %s", src, dest)

	return retVal
}

func Which(c *Command) Result {

	args, r := GetStringArrayArgument(c, "args")
	if r != nil {
		return *r
	}

	retVal := Result{Type: OK, Info: make(map[string]interface{}, 0)}

	for _, name := range args {
		path, ok := internal.ExistsInPath(name)
		if !ok {
			retVal.Error = errors.New("not found in PATH")
			retVal.Type = ERROR
			retVal.Info[name] = "NOT FOUND IN PATH"
		}
		retVal.Info[name] = path
	}

	return retVal
}

func Find(c *Command) Result {

	root, r := GetStringArgument(c, "path")
	if r != nil {
		return *r
	}
	pattern, r := GetStringArgument(c, "pattern")
	if r != nil {
		return *r
	}

	retVal := Result{Type: OK}
	foundList := make(map[string]interface{})

	if re, e := regexp.Compile(pattern); e == nil {
		e = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return filepath.SkipDir
			}

			if re.MatchString(path) {
				foundList[path] = fmt.Sprintf("%s %d", info.Mode().String(), info.Size())
			}
			return nil
		})
	} else {
		return Result{Type: ERROR, Error: e}
	}

	retVal.Info = foundList

	return retVal
}
