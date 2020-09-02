package v1

import (
  "strings"
)

// Is a command result and contains various fields
type Result struct {
  Type    ResultType
  Error   error
  Name    string
  Message string
  Info    map[string]interface{}
}

type ResultType string

const (
  CHANGED ResultType = "CHANGED"
  OK      ResultType = "OK"
  ERROR   ResultType = "ERROR"
  SKIPPED ResultType = "SKIPPED"
)

type ArgumentOption struct {
  Required   bool
  Default    interface{}
  HelpString string
}

type Command struct {
  // Name of the command
  Name string
  // Various command aliases
  Aliases []string
  // Function to call when invoked
  Run func(*Command) Result
  // Help string
  Help string
  // Arguments help string
  Args    map[string]string


  // Vars injected from executor
  vars map[string]interface{}
  // command arguments from file
  args map[string]interface{}
  // output callback from executor
  callback OutputCallback
}

var DefaultCommandList = []Command{
  {
    Name:    "unzip",
    Aliases: []string{"archive"},
    Help:    "Extracts zip archives in a given folder",
    Args: map[string]string{
      "src":  "Zip file",
      "dest": "Destination Folder",
    },
    Run: Unzip,
  },
  {
    Name:    "brew",
    Aliases: []string{"homebrew"},
    Help:    "Handle homebrew actions",
    Args: map[string]string{
      "cmd": "brew action (any valid brew subcommand)",
      "args": "A list of brew arguments",
    },
    Run: Brew,
  },
  {
    Name:    "hashfile",
    Aliases: []string{"sha256file"},
    Help:    "Calculate SHA256 hash of a file",
    Args: map[string]string{
      "args":  "List of file path",
    },
    Run: Hashfile,
  },
  {
    Name:    "hash",
    Aliases: []string{"sha256"},
    Help:    "Calculate SHA256 hash of a string",
    Args: map[string]string{
      "args":  "List of file path",
    },
    Run: HashString,
  },
  {
    Name:    "download",
    Aliases: []string{"get_uri", "dwn"},
    Help:    "Download data from remote URI saving it to dest file",
    Args: map[string]string{
      "src":  "The source URI",
      "dest": "Destination File",
      "mode": "File creation mode",
    },
    Run: Download,
  },
  {
    Name:    "echo",
    Aliases: []string{"Echo", "print"},
    Help:    "Print out the give strings",
    Args: map[string]string{
      "args": "A list of strings",
    },
    Run: Echo,
  },
  {
    Name:    "exec",
    Aliases: []string{"Exec"},
    Help:    "Execute a binary with arguments",
    Args: map[string]string{
      "cmd":  "Command name (path of binary)",
      "args": "argument list",
    },
    Run: Exec,
  },
  {
    Name:    "shell",
    Aliases: []string{"Shell", "sh", "bash"},
    Help:    "Execute commands using default shell",
    Args: map[string]string{
      "args": "A list of arguments",
      "cwd":  "change working directory",
    },
    Run: Shell,
  },
  {
    Name:    "copy",
    Aliases: []string{"Copy", "cp"},
    Help:    "Copy source path to dest",
    Args: map[string]string{
      "src":  "The source path",
      "dest": "Destination path",
    },
    Run: Copy,
  },
  {
    Name:    "exists",
    Aliases: []string{"exist", "is"},
    Help:    "Return OK if the path currently exists",
    Args: map[string]string{
      "args": "A list of paths",
    },
    Run: FileExists,
  },
  {
    Name:    "link",
    Aliases: []string{"Link", "ln"},
    Help:    "Create a symbolic link like the ln command",
    Args: map[string]string{
      "src":  "The source path",
      "dest": "Destination path",
    },
    Run: Link,
  },
  {
    Name:    "mkdir",
    Aliases: []string{"Mkdir", "makedir"},
    Help:    "Create a directory",
    Args: map[string]string{
      "args": "A list of paths",
    },
    Run: Mkdir,
  },
  {
    Name:    "move",
    Aliases: []string{"Move", "mv", "rename"},
    Help:    "Rename a path",
    Args: map[string]string{
      "src":  "Source path",
      "dest": "Destination path",
    },
    Run: Move,
  },
  {
    Name:    "remove",
    Aliases: []string{"Remove", "rm"},
    Help:    "Remove one or more file or directories",
    Args: map[string]string{
      "args": "A list of paths",
    },
    Run: RemoveFiles,
  },
  {
    Name:    "which",
    Aliases: []string{"Which"},
    Help:    "Return the current absolute path of a given executable if in PATH",
    Args: map[string]string{
      "args": "A list of binary names",
    },
    Run: Which,
  },
  {
    Name:    "find",
    Aliases: []string{"Find", "search"},
    Help:    "Looks for specified patterns in given directory",
    Args: map[string]string{
      "path":    "The root directory where start search",
      "pattern": "Search pattern",
    },
    Run: Find,
  },
  {
    Name:    "template",
    Aliases: []string{"Template", "tpl"},
    Help:    "Use the given source file as template, resolve variables and output it to dest",
    Args: map[string]string{
      "src":  "Template file path",
      "dest": "Destination file path",
    },
    Run: Template,
  },
  {
    Name:    "setvars",
    Aliases: []string{"set"},
    Help:    "Set variables",
    Args:    map[string]string{},
    Run:     SetVars,
  },
  {
    Name:    "printvars",
    Aliases: []string{"debugvars", ""},
    Help:    "Print out all recorded vars",
    Args:    map[string]string{},
    Run:     PrintVars,
  },
}

// Return boolean when name given match Name or Aliases
func (c Command) Is(name string) bool {
  if strings.EqualFold(c.Name, name) {
    return true
  }

  for _, alas := range c.Aliases {
    if strings.EqualFold(alas, name) {
      return true
    }
  }
  return false
}

// Execute the Run function only if name match Name or Aliases
func (c Command) Execute(args map[string]interface{}, vars map[string]interface{}, outCallback OutputCallback) Result {
  c.args = args
  c.vars = vars
  c.callback = outCallback
  return c.Run(&c)
}
