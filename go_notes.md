# Go notes.

- Go code in single workspace

- workspace contains vcs repos

- repo has >= 1 packages

- path to pkg dir determines import path

## Workspace

Directory hierarchy w/ 3 directories at root

- src # Go source files

	- vcs goes here, tracks dev of source code

- pkg # Contains package objects

- bin # contains executable commands

## GOPATH

- specifies location of the workspace.

defaults to home_directory/go

set GOPATH to location of workspace

https://github.com/golang/go/wiki/SettingGOPATH

export GOPATH=$HOME/Go

## Important Paths

String that uniquely identifies a package.

Package's important path == location inside a workspace or remote repo

### Standard libs

given short important paths i.e "fmt", "net/http"

### Own packages 

Choose a base path that won't collide with future additions to std lib or other external libs.

Base path = root of source repo

## Building and Installing

path structure needs to be Go/src/<foldername>/has.go file

$ go install [don't put #GOPATH/src/]hello

or 

$ cd $GOPATH/src/<path to stuff>/stuff # where stuff has .go files.

builds hello command producing an executable binary. Then installs binary to workspace's bin directory as hello.

## Running program

$ $GOPATH/bin/hello

or 

$ cd $GOPATH/bin

$ ./hello

```none
Note: 

- Only install produces the binary executable, build just builds it.

- if installing something that imports something, it installs dependencies as well
```

## Packages 

- Installing a library puts a lib_name.a file in dir pkg/os_architecture/

- Statically linked, pkg objects don't need to be present to run Go programs.

first statement of a Go source file must be 
package name

### Libraries --> package <lib_name>
### Main file/runnable --> package main (Executable commands in main)

Files in same package use same name

Naming convention, package name == last element of import path

package imported as "pokemon/faves/elec/magnemite" named "magnemite"

No requirement package names are unique across all packages linked to a single binary, only that import paths are unique.

