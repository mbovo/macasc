package internal

import (
  "os"
  "os/exec"
)

func Exec(silent bool, cmd string, arg ...string) (int, error) {

  var args = []string{cmd}
  args = append(args, arg...)

  c := exec.Command(cmd)
  c.Args = args
  if !silent {
    c.Stdout = os.Stdout
    c.Stdin = os.Stdin
    c.Stderr = os.Stderr
  }
  err := c.Run()
  return c.ProcessState.ExitCode(), err
}

func Shell(chdir string, arg ...string) (int, error) {
  cwd, _ := os.Getwd()
  if chdir != "" {
    e := os.Chdir(chdir)
    if e != nil {
      return 0, e
    }
  }

  var args = []string{"-c"}
  args = append(args, arg...)
  i, e := Exec(false, "/bin/bash", args...)
  os.Chdir(cwd)
  return i, e
}
