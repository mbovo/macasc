package v1

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
)

type DefaultCallback struct {
	stdout *os.File
	stderr *os.File
}

func NewDefaultCallback(out, err *os.File) OutputCallback {

	if out == nil {
		out = os.Stdout
	}
	if err == nil {
		err = os.Stderr
	}

	return DefaultCallback{
		stdout: out,
		stderr: err,
	}
}

func (d DefaultCallback) Output(format string, a ...interface{}) {
	_, err := fmt.Fprintf(d.stdout, format, a...)
	if err != nil {
		panic(err)
	}
}

func (d DefaultCallback) Error(format string, a ...interface{}) {
	_, err := fmt.Fprintf(d.stderr, color.Error.Render(format), a...)
	if err != nil {
		panic(err)
	}
}

func (d DefaultCallback) Header(format string, a ...interface{}) {
	d.Output(format, color.Note.Render(a...))
}

func (d DefaultCallback) Result(r Result) {

	c := color.FgDefault.Render
	switch r.Type {
	case ERROR:
		c = color.Style{color.Red, color.OpBold}.Render
	case OK:
		c = color.Style{color.Blue, color.OpBold}.Render
	case CHANGED:
		c = color.Style{color.Yellow, color.OpBold}.Render
	case SKIPPED:
		c = color.Style{color.Cyan, color.OpBold}.Render
	}

	substr := strings.ReplaceAll(r.Message, "\n", "\\n")
	d.Output("[%7s]  %s\n", c(r.Type), color.Secondary.Render(substr))

	if r.Error != nil {
		d.Error("%20s %s\n", "", r.Error)
	}

	if r.Info != nil {
		for key, item := range r.Info {
			d.Output("%20s %s\t%s\n", "", color.Secondary.Render(key), color.Secondary.Render(item))
		}
	}
}
