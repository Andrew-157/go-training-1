# Go Tutorial

- [Go Tutorial](#go-tutorial)
  - [Introduction](#introduction)
  - [Use custom Go module without publishing it](#use-custom-go-module-without-publishing-it)

## Introduction

This document describes different useful stuff about `Go` while I am learning it

## Use custom Go module without publishing it

The following steps will allow you to use a module on a local filesystem without publishing it:

```commandline
root@fedora:~/go_dir# mkdir helloworld
root@fedora:~/go_dir# cd helloworld/
root@fedora:~/go_dir/helloworld# go mod init example.com/helloworld
go: creating new go.mod: module example.com/helloworld
root@fedora:~/go_dir/helloworld# ls
go.mod
root@fedora:~/go_dir/helloworld# touch helloworld.go
root@fedora:~/go_dir/helloworld# ls
go.mod  helloworld.go
root@fedora:~/go_dir/helloworld# cat helloworld.go 
/*
Simple package that has one function to print "Hello World!"
*/
package helloworld

import "fmt"

// Prints "Hello World!"
func HelloWorld() {
        fmt.Println("Hello World!")
}
root@fedora:~/go_dir/helloworld# cd ..
root@fedora:~/go_dir# mkdir app
root@fedora:~/go_dir# cd app
root@fedora:~/go_dir/app# go mod init example.com/app
go: creating new go.mod: module example.com/app
root@fedora:~/go_dir/app# ls
go.mod
root@fedora:~/go_dir/app# cat go.mod
module example.com/app

go 1.23.4
root@fedora:~/go_dir/app# touch run.go
root@fedora:~/go_dir/app# ls
go.mod  run.go
root@fedora:~/go_dir/app# cat run.go 
package main

import (
        "example.com/helloworld"
)

func main() {
        helloworld.HelloWorld()
}
root@fedora:~/go_dir/app# go mod edit -replace example.com/helloworld=../helloworld/
root@fedora:~/go_dir/app# cat go.mod 
module example.com/app

go 1.23.4

replace example.com/helloworld => ../helloworld/
root@fedora:~/go_dir/app# go mod tidy
go: found example.com/helloworld in example.com/helloworld v0.0.0-00010101000000-000000000000
root@fedora:~/go_dir/app# cat go.mod
module example.com/app

go 1.23.4

replace example.com/helloworld => ../helloworld/

require example.com/helloworld v0.0.0-00010101000000-000000000000
root@fedora:~/go_dir/app# go run .
Hello World!
```
