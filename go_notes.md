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


