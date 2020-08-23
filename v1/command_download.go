package v1

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/mbovo/yacasc/v1/internal"
)

// Download command will try to download the src to the dest file
// Accepted arguments:
//	src: string required
//	dest: string required
//	mode: uint32 optional dest fileMode
func Download(c *Command) Result {

	if e := internal.VerifyRequiredArgs(c.Name, []string{"src","dest"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	src := c.args["src"]
	dest := c.args["dest"]
	mode, ok := c.args["mode"]
	if !ok {
		mode = os.ModePerm
	}

	retVal := Result{Type: OK}

	info, err := os.Stat(dest.(string))
	if err == nil {
		// file exists
		if info.IsDir() {
			// target exist and is a directory, append the file name and use it as destination
			dest = filepath.Join(dest.(string), path.Base(src.(string)))
		}
	} else {
		if os.IsNotExist(err) {
			// ok file does not exists, start download
			if e := internal.DownloadToFile(src.(string), dest.(string), mode.(os.FileMode)); e != nil {
				retVal.Type = ERROR
				retVal.Error = e
			}
			retVal.Type = CHANGED
			retVal.Message = fmt.Sprintf("Downloaded %s -> %s", src, dest)
		} else {
			retVal.Type = ERROR
			retVal.Error = err
		}
	}
	return retVal
}
