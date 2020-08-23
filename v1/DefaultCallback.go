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
	_, err := fmt.Fprintf(d.stderr, format, color.Error.Render(a...))
	os.Exit(1)
	if err != nil {
		panic(err)
	}
}

func (d DefaultCallback) Header(format string, a ...interface{}){
	d.Output(format, color.Note.Render(a...))
}

func (d DefaultCallback) Result(r Result) {

	red := color.Style{color.Red, color.OpBold}.Render
	green := color.Style{color.Green, color.OpBold}.Render
	yellow := color.Style{color.Yellow, color.OpBold}.Render
	c := color.FgDefault.Render

	substr := strings.ReplaceAll(r.Message, "\n", "\\n")

	switch r.Type {
	case ERROR:
		c = red
	case OK:
		c = green
	case CHANGED:
		c = yellow
	}

	d.Output("[%7s]  %s\n", c(r.Type), color.Secondary.Render(substr))

	if r.Error != nil {
		d.Error("%20s %s\n", "", r.Error)
	}
	//if r.Message != "" {
	//	d.Output("\t%s\n", r.Message)
	//}
	if r.Info != nil {
		for key, item := range r.Info {
			d.Output("%20s %s\t%s\n", "", color.Secondary.Render(key), color.Secondary.Render(item))
		}
	}
}
