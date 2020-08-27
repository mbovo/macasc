package internal

import (
  "os"

  "gopkg.in/flosch/pongo2.v3"
)

func Template(s string, context pongo2.Context) (string, error) {
  tpl, e := pongo2.FromString(s)
  if e != nil {
    return "", e
  }
  s, e = tpl.Execute(context)
  if e != nil {
    return "", e
  }
  return s, nil
}

func TemplateFile(srcFile string, destFile string, context pongo2.Context) error {

  tpl, err := pongo2.FromFile(srcFile)
  if err != nil {
    return err
  }
  out, err := os.Create(destFile)
  if err != nil {
    return err
  }
  defer out.Close()

  err = tpl.ExecuteWriter(context, out)
  if err != nil {
    return err
  }
  return nil
}
