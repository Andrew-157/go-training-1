# Go Tutorial

1. [Introduction](#introduction)
2. [Use custom Go module without publishing it](#use-custom-go-module-without-publishing-it)
3. [What is a Pointer](#what-is-a-pointer)
4. [Nested Unpublished Packages](#nested-unpublished-packages)

## Introduction

This document describes different useful stuff about `Go` while I am learning it

## Use custom Go module without publishing it

> `NOTE`: I took it from here: [Call your code from another module](https://go.dev/doc/tutorial/call-module-code)

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

## What is a Pointer

> `WARNING`: For now it is very short section, so that I can quickly remember what pointer is and how it works.

A variable is a piece of storage containing a value. Variables created by declarations are identified by a name, such as `x`, but many variables are identified only by expressions like `x[i]` or `x.f`. All these expressions read the value of a variable, except when they appear on the left hand side of an assignment, in which case a new value is assigned to the variable.

A pointer value is the address of a variable. A pointer is thus the location at which a value is
stored. Not every value has an address, but every variable does. With a pointer, we can read or update the value of a variable indirectly, without using or even knowing the name of the variable, if indeed it has a name.

If a variable is declared `var x int`, the expression `&x` ("address of x") yields a pointer to an integer variable, that is, a value of type `*int`, which is pronounced "pointer to int." If this value is called `p`, we say "p points to x," or equivalently "p contains the address of x." The variable to which `p` points is written `*p`. The expression `*p` yields the value of that variable, an `int`, but since `*p` denotes a variable, it may also appear on the left-hand side of an assignment, in which case the assignment updates the variable.

> `NOTE`: I took this explanation from this book: ["The Go Programming Language"](https://github.com/neo-liang-sap/book/blob/master/Go/The.Go.Programming.Language.pdf)

Example:

```go
x := 1
var p *int = &x                        // p, of type *int points, to x
fmt.Printf("Address of x is %v\n", p)  // Output: Address of x is 0xc000012110
fmt.Printf("x == *p is %v\n", x == *p) // Output: x == *p is true
*p = 45                                // Equivalent to x = 45
fmt.Printf("Now x is %d\n", x)         // Output: Now x is 45
```

## Nested Unpublished Packages

See [document](./nested-unpublished-packages/Nested-unpublished-packages.md)
