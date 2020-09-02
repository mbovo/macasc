package main

import (
  "fmt"
  "log"
  "os"
  "path/filepath"
  "runtime"
  "strings"

  "github.com/spf13/cobra"
  "gopkg.in/yaml.v2"

  v1 "github.com/mbovo/yacasc/v1"
)


var Version string
var Build string
var cfgFile, varFile string

var rootCmd = &cobra.Command{
  Use:   "yacasc",
  Short: "Configuration As Code",
  Long:  `Yet Another Configuration as Code`,
}

func main() {
  Execute()
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    log.Fatal(err)
  }
}

func init() {
  cobra.OnInitialize(initConfig)
  rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "","Global configuration file")
  rootCmd.PersistentFlags().StringVarP(&varFile, "vars", "v", "","Variables file")
  log.SetOutput(os.Stderr)
  rootCmd.AddCommand(helpCmd)
  rootCmd.AddCommand(verifyCmd)
  rootCmd.AddCommand(runCmd)
  rootCmd.AddCommand(printCmd)
  rootCmd.AddCommand(listCmd)
  rootCmd.AddCommand(versionCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile == "" {
    cfgFile = filepath.Join(".", "config.yml")
  }
}

func run(cmd *cobra.Command, args []string) {
  if len(args) < 1 {
    log.Fatal("Empty stepFile, please provide a valid filename or uri")
  }

  var e error
  executor, e := v1.NewExecutor(v1.NewDefaultCallback(nil, nil) , args[0], varFile)

  if e != nil {
    log.Fatal(e)
  }

  switch cmd.Use {
  case "run":
    e = executor.Run()
    break
  case "print":
    y, _ := yaml.Marshal(executor.Vars)
    fmt.Printf("%s\n", y)
    y, _ = yaml.Marshal(executor.Steps)
    fmt.Printf("#############################################################################\n%s\n", y)
  case "verify":
    log.Printf("Not Yet implemented")
  default:
    log.Printf("Unknown command %s\n", cmd.Use)
    break
  }

  if e != nil {
    log.Fatal(e)
  }

}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print version and exit",
  Long:  `Print version and exit`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("yacasc v%s.%s - %s.%s\n", Version, Build, runtime.GOOS, runtime.GOARCH)
    os.Exit(0)
  },
}

var listCmd = &cobra.Command{
  Use:   "list",
  Short: "List available command",
  Long:  `List all available command`,
  Run: listCommands,
}

var helpCmd = & cobra.Command{
  Use: "command",
  Short: "Show command usage",
  Long: `Show command usage`,
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string){
    cmdName := args[0]
    for _, comm := range v1.DefaultList {
      if strings.EqualFold(cmdName, comm.Name){
        fmt.Fprintf(os.Stderr,"[%s]\t\t%s\n", comm.Name, comm.Help)
        for argName, argHelp := range comm.Args {
          fmt.Fprintf(os.Stderr,"  %s\t%s\n", argName, argHelp)
        }
        os.Exit(0)
      }
    }
    fmt.Fprintf(os.Stderr, "ERR: command %s not found\n",cmdName)
    listCommands(cmd, args)
  },
}

func listCommands (cmd *cobra.Command, args []string) {
  cmds := v1.DefaultList
  fmt.Fprintln(os.Stderr,"Available Commands:")
  for _, cmd := range cmds {
    fmt.Fprintf(os.Stderr,"  %10s\t%s\n", cmd.Name, cmd.Help)
  }
}

var verifyCmd = &cobra.Command{
  Use:   "verify",
  Short: "Try to identify changes",
  Long:  `Plan the execution identifying changes`,
  Run:   run,
}


var runCmd = &cobra.Command{
  Use:   "run",
  Short: "Run the step file ",
  Long:  `Execute all steps and commands in the step file given as argument`,
  Run:   run,
}

var printCmd = &cobra.Command{
  Use:   "print",
  Short: "Print internal info",
  Long:  `Print internal information regarding loaded vars, steps, and commands`,
  Run:   run,
}