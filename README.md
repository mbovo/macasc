# Yet Another Configurator As Code

![Go](https://github.com/mbovo/yacasc/workflows/Go/badge.svg?branch=master)

A little golang tool that try to implement "as code" approach 
 
```shell script
yacasc v0.0.3.b42d1a3 - linux.amd64
Yet Another Configuration as Code

Usage:
  yacasc [command]

Available Commands:
  add         Add 
  apply       Apply the configuration
  command     Show command usage
  configure   Configuration step
  help        Help about any command
  list        List available command
  remove      Remove command
  verify      Try to identify changes
  version     Print version and exit

Flags:
  -c, --config string   Global configuration file
  -h, --help            help for yacasc
  -v, --vars string     Variables file

Use "yacasc [command] --help" for more information about a command.

```

## Implemented commands

```shell script
$ yacasc list
yacasc v0.0.3.b42d1a3 - linux.amd64
Available Commands:
       unzip	Extracts zip archives in a given folder
        brew	Handle homebrew actions
    hashfile	Calculate SHA256 hash of a file
        hash	Calculate SHA256 hash of a string
    download	Download data from remote URI saving it to dest file
        echo	Print out the give strings
        exec	Execute a binary with arguments
       shell	Execute commands using default shell
        copy	Copy source path to dest
      exists	Return OK if the path currently exists
        link	Create a symbolic link like the ln command
       mkdir	Create a directory
        move	Rename a path
      remove	Remove one or more file or directories
       which	Return the current absolute path of a given executable if in PATH
    template	Use the given source file as template, resolve variables and output it to dest
     setvars	Set variables
   printvars	Print out all recorded vars
```

## Example of step file

```yaml
---
# This is list of steps

# This is a step called "test"
- name: test
  # each step can have 4 different type of commands, they are tight with
  #   cli commands: add , remove, configure, verify
  add: []
  remove: []
  configure: []
  verify:
    # This is a list of commands
    # using `yacasc list` you can output the whole list of commands available
    # other keys except "name" are loaded by the command and the syntax depends on the cmd itself
    - name: echo
      args:
        # you can use variables defined in the var file, the syntax is Jinja2 compliant
        - {{ hallostring }}
  {% for msg in messages %}
        - {{ msg }}
  {% endfor %}
    - name: download
      src: {{brew.url}}
      dest: {{brew.tmpfile}}

    - name: download
      src: {{brew.url}}
      dest: {{brew.tmpfile}}

    - name: download
      src: {{vscode.url}}
      dest: {{vscode.tmpfile}}

    - name: unzip
      src: {{vscode.tmpfile}}
      dest: /tmp/

    - name: hashfile
      args:
        - {{brew.tmpfile}}

    - name: hash
      args: ["myname"]

    - name: copy
      src: {{ brew.tmpfile }}
      dest: /tmp/anotherone

    - name: copy
      src: {{ brew.tmpfile }}
      dest: /tmp/anotherone

    - name: copy
      src: {{ brew.tmpfile }}
      dest: /tmp/anotherone
      force: true

    - name: copy
      src: /not/found
      dest: /tmp/anotherone

    - name: mkdir
      args:
        - /tmp/test

    - name: remove
      args:
        - {{brew.tmpfile}}
        - {{vscode.tmpfile}}
        - /tmp/test
        - /tmp/test
        - /tmp/anotherone
        - /tmp/Visual Studio Code.app
```
