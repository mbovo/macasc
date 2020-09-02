# Yet Another Configurator As Code

![Go](https://github.com/mbovo/yacasc/workflows/Go/badge.svg)

A little golang tool that try to implement "as code" approach 
 
```
yacasc v0.0.2.12083d9 - darwin.amd64
Yet Another Configuration as Code

Usage:
  yacasc [command]

Available Commands:
  command     Show command usage
  help        Help about any command
  list        List available command
  print       Print internal info
  run         Run the step file 
  verify      Try to identify changes
  version     Print version and exit

Flags:
  -c, --config string   Global configuration file
  -h, --help            help for yacasc
  -v, --vars string     Variables file

Use "yacasc [command] --help" for more information about a command.


```

## Implemented commands

```
$ yacasc list
yacasc v0.0.2.12083d9 - darwin.amd64
Available Commands:
       unzip    Extracts zip archives in a given folder
        brew    Handle homebrew actions
    hashfile    Calculate SHA256 hash of a file
        hash    Calculate SHA256 hash of a string
    download    Download data from remote URI saving it to dest file
        echo    Print out the give strings
        exec    Execute a binary with arguments
       shell    Execute commands using default shell
        copy    Copy source path to dest
      exists    Return OK if the path currently exists
        link    Create a symbolic link like the ln command
       mkdir    Create a directory
        move    Rename a path
      remove    Remove one or more file or directories
       which    Return the current absolute path of a given executable if in PATH
        find    Looks for specified patterns in given directory
    template    Use the given source file as template, resolve variables and output it to dest
     setvars    Set variables
   printvars    Print out all recorded vars
```

## Example of step file

```yaml
---
- name: StepName
  vars:
    sayHallo: false
  cmds:
    - echo:
        msg: message in the bottle
    - echo:
        msg: hallo world
      opts:
        when: "{{sayHallo}}"
    - find:
        path: /tmp/
        pattern: "file"
      opts:
        register: fileList
    - hashFile:
        args: []
        fromVar: fileList
      opts:
        register: myvar
    - printvars:
```
The output will be (assuming a `/tmp/file` exists):
```
yacasc v0.0.2.12083d9 - darwin.amd64
Loaded resources/stepv2.yml 
Step 1/1: Step1
-       echo  [OK]  message in the bottle\n
-       echo  [SKIPPED]  
-       find  [OK]  
                     /tmp/file  -rw-r--r-- 75
-   hashFile  [OK]  
                     /tmp/file  8503ff8fcb48ecf19620378aa2027f4a3045674f8a621d8239f681ea0a1b0423
-  printvars  [OK]  
                     myvar      map[/tmp/file:8503ff8fcb48ecf19620378aa2027f4a3045674f8a621d8239f681ea0a1b0423]
                     sayHallo   false
                     fileList   map[/tmp/file:-rw-r--r-- 75]

```
