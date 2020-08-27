package v1

import (
  "fmt"
  "os"
  "path"
  "path/filepath"
  "strconv"

  "github.com/mbovo/yacasc/v1/internal"
)

// Download command will try to download the src to the dest file
// Accepted arguments:
//	src: string required
//	dest: string required
//	mode: uint32 optional dest fileMode
func Download(c *Command) Result {

  src, r := GetStringArgument(c, "src")
  if r != nil { return *r }
  dest, r := GetStringArgument(c, "dest")
  if r != nil { return *r }

  var mode os.FileMode
  m, r := GetStringArgument(c, "mode")
  if r != nil {
    mode = os.ModePerm
  }else{
    i, _ := strconv.Atoi(m)
    mode = os.FileMode(i)
  }

  retVal := Result{Type: OK}

  info, err := os.Stat(dest)
  if err == nil {
    // file exists
    if info.IsDir() {
      // target exist and is a directory, append the file name and use it as destination
      dest = filepath.Join(dest, path.Base(src))
    }
  } else {
    if os.IsNotExist(err) {
      // ok file does not exists, start download
      if e := internal.DownloadToFile(src, dest, mode); e != nil {
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
