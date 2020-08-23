package v1

import (
	"strings"

	"github.com/mbovo/yacasc/v1/internal"
)


func Brew(c *Command) Result {

	if e := internal.VerifyRequiredArgs(c.Name,[]string{"args", "cmd"}, c.args); e != nil {
		return Result{Type: ERROR, Error: e}
	}

	argsInt := c.args["args"].([]interface{})

	brewArgs := make([]string,0)
	for _, i := range argsInt{
		brewArgs = append(brewArgs, i.(string))
	}

	cmd := c.args["cmd"].(string)

	retVal := Result{Type: CHANGED}

	//TODO: return stderr and stdout content and write it to retVal.Info
	if err := BrewAction(strings.ToLower(cmd), brewArgs, c.callback); err != nil {
		retVal.Error = err
		retVal.Type = ERROR
	}

	return retVal
}

func BrewAction(command string, list []string, callback OutputCallback ) error {

	brewBin, ok := internal.ExistsInPath("brew")
	if !ok {
		//trying without abs location
		brewBin = "brew"
	}

	for _, formula := range list {
		callback.Output("%s", formula)
		if _, err := internal.Exec(true, brewBin, command, formula ); err != nil {
			return err
		}
	}
	return nil
}
//
//func preFetch(formule []string) (int, error) {
//	s := []string{"fetch"}
//	if len(formule) < 1 {
//		return 0, nil
//	}
//
//	s = append(s, formule...)
//
//	log.Printf("+ Pre-fetching %d formulas\n", len(formule))
//	return internal.Exec(true, "/usr/local/bin/brew", s...)
//}
//
//func filterFormulae(formule []string) []string {
//	log.Printf("+ Check %d formulae\t", len(formule))
//	c := make(chan string, len(formule)+1)
//	for _, f := range formule {
//		go listFormula(f, c)
//	}
//	var l []string
//	for range formule {
//		s := <-c
//		if s != "" {
//			l = append(l, s)
//		}
//	}
//	log.Printf("%d missing\n", len(l))
//	return l
//}
//
//func listFormula(formula string, c chan string) {
//	rc, err := internal.Exec(true, "/usr/local/bin/brew", ListCommand, formula)
//	if rc != 0 || err != nil {
//		c <- formula
//	}
//	c <- ""
//}
