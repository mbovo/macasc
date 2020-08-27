package internal

import "testing"

func TestExec(t *testing.T) {
  mustOK, err1 := Exec(true, "/bin/bash", "-c", "echo", "hallo")
  _, err2 := Exec(true, "/bin/sdfnauhwer")
  mustKO, err3 := Exec(true, "/bin/bash", "-c", "false")
  mustOK2, err4 := Exec(false, "/bin/bash", "-c", "cat", "/etc/hosts")
  if mustOK != 0 || err1 != nil {
    t.Errorf("Error: /bin/bash -> %v %v", mustOK, err1)
  }
  if err2 == nil {
    t.Errorf("Error: /bin/sdfnauhwer -> %v", err2)
  }
  if mustKO == 0 || err3 == nil {
    t.Errorf("Error: /bin/bash -c false -> %v %v", mustKO, err3)
  }
  if mustOK2 != 0 || err4 != nil {
    t.Errorf("Error: /bin/bash -c cat /etc/hosts %v %v", mustKO, err3)
  }
}

func TestShell( t *testing.T){
  ok, err := Shell("", "cat", "/etc/hosts")

  if ok != 0 || err != nil {
    t.Errorf("Error executing [cat /etc/hosts] in ''")
  }

  ok, err = Shell("/wrong/dir", "cat", "/etc/hosts")

  if ok != 0 || err == nil {
    t.Errorf("Error: cwd to /wrong/dir must fail")
  }

  ok, err = Shell("/etc", "cat", "hosts")
  if ok != 0 || err != nil {
    t.Errorf("Error executing [cat hosts] in /etc")
  }

  _, err = Shell ("", "booooo")

  if err == nil {
    t.Errorf("Error: execution of [boooo] should fail")
  }

}