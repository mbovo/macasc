package v1

import (
	"os"
	"testing"
)

func TestNewDefaultCallback(t *testing.T){

	table := []struct {
		p1 *os.File
		p2 *os.File
		res DefaultCallback
	} {
		{nil,nil, DefaultCallback{
			stdout: os.Stdout,
			stderr: os.Stderr,
		}},
		{ os.Stderr, os.Stdout, DefaultCallback{
			stdout: os.Stderr,
			stderr: os.Stdout,
		}},
	}

	for _, test := range table {
		cb := NewDefaultCallback(test.p1, test.p2)
		if cb != test.res {
			t.Errorf("Want %#v got %#v", cb, test.res)
		}
	}



}