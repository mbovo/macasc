package internal

import (
  "archive/zip"
  "crypto/sha256"
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "path/filepath"
  "strings"
)

func Unzip(src string, dest string) ([]string, error) {
  var files []string

  r, err := zip.OpenReader(src)
  if err != nil {
    return files, err
  }
  defer r.Close()

  for _, f := range r.File {

    fullPath := filepath.Join(dest, f.Name)

    // check for zipslip vuln https://snyk.io/research/zip-slip-vulnerability#go
    if !strings.HasPrefix(fullPath, filepath.Clean(dest)+string(os.PathSeparator)) {
      return nil, fmt.Errorf("illegal file path %s", fullPath)
    }

    files = append(files, fullPath)

    // make folder if any
    if f.FileInfo().IsDir() {
      if e := os.Mkdir(fullPath, os.ModePerm); e != nil {
        return nil, e
      }
      continue
    }

    // make destination path
    if e := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); e != nil {
      return nil, e
    }

    // open destination file
    oFile, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
    if err != nil {
      return nil, err
    }

    // open input file
    iFile, err := f.Open()
    if err != nil {
      return nil, err
    }

    // copy content
    _, err = io.Copy(oFile, iFile)

    oFile.Close()
    iFile.Close()

    if err != nil {
      return nil, err
    }

  }

  return files, nil
}

func CopyFile(src string, dest string) error {

  in, err := os.Open(src)
  if err != nil {
    return err
  }
  out, err := os.Create(dest)
  if err != nil {
    return err
  }

  defer in.Close()
  defer out.Close()

  written, err := io.Copy(out, in)
  if err != nil {
    return err
  }

  stat, _ := in.Stat()
  if stat.Size() != written {
    return fmt.Errorf("error copying, written %d/%d bytes", written, stat.Size())
  }

  return nil
}

func DownloadToFile(url string, filepath string, mode os.FileMode) error {

  // Get the data
  if resp, err := http.Get(url); err == nil {
    defer resp.Body.Close()
    if out, err := os.Create(filepath); err == nil {
      defer out.Close()
      if _, err := io.Copy(out, resp.Body); err == nil {
        if err := os.Chmod(filepath, mode); err != nil {
          return err
        }
      } else {
        return err
      }
    } else {
      return err
    }
  } else {
    return err
  }
  return nil
}

func Download(url string) ([]byte, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  content, err := ioutil.ReadAll(resp.Body)

  return content, err
}

func FileExists(filename string) bool {
  info, err := os.Stat(filename)
  if os.IsNotExist(err) {
    return false
  }
  return !info.IsDir()
}

// Return the absolute path of given filename and optionally boolean false when not found
func ExistsInPath(filename string) (string, bool) {
  for _, v := range strings.Split(os.Getenv("PATH"), ":") {
    match, err := filepath.Glob(filepath.Join(v, filename))
    if err == nil && len(match) > 0 {
      return match[0], true
    }
  }
  return "",false
}

func MakeDir(path string, perms os.FileMode) error {
  return os.MkdirAll(path, perms)
}

func TouchFile(path string) (*os.File, error) {
  return os.Create(path)
}

func LoadFromFileOrUri(uri string) (content []byte, e error) {

  if strings.HasPrefix(uri, "http") {
    content, e = Download(uri)
    if e != nil {
      return content, e
    }
  } else {
    if !FileExists(uri) {
      return content, fmt.Errorf("file not found %s\n", uri)
    }
    content, e = ioutil.ReadFile(uri)
  }
  return content, e
}

func SHA256String(s string) string {
  sum := sha256.Sum256([]byte(s))
  return fmt.Sprintf("%x", sum)
}

func SHA256File(path string) (string, error) {
  if !FileExists(path) {
    return "", fmt.Errorf("file not found")
  }

  f, err := os.Open(path)
  if err != nil {
    return "", err
  }
  defer f.Close()

  h := sha256.New()
  if _, err := io.Copy(h, f); err != nil {
    return "", err
  }
  return fmt.Sprintf("%x", h.Sum(nil)), nil
}
