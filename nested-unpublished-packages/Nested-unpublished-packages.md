# How to use Nested Unpublished Packages

## Set up some package

In this directory you will find package `utils`, that is used by the upper level package `server`

>`NOTE`: `utils` package and all others described here are unpublished.

## Set up package which will use package from the previous step

In this directory you will find package `server` that uses package `utils` from the previous step.

>`NOTE`: To use unpublished package from local filesystem [see](../README.md#use-custom-go-module-without-publishing-it).

## Use an unpublished package that uses other unpublished package

In this directory, you will find package `app` that directly uses `server` package and indirectly uses `utils` package.

Here is how its `go.mod` file looks like:

```mod
module example.com/app

go 1.23.4

replace example.com/server => ../server/

replace example.com/utils => ../utils/

require example.com/server v0.0.0-00010101000000-000000000000

require example.com/utils v0.0.0-00010101000000-000000000000 // indirect

```

Here are the steps:

```commandline
root@fedora:~/go_dir# mkdir app
root@fedora:~/go_dir# cd app
root@fedora:~/go_dir/app# touch run.go
root@fedora:~/go_dir/app# go mod init example.com/app
go: creating new go.mod: module example.com/app
go: to add module requirements and sums:
        go mod tidy
root@fedora:~/go_dir/app# cat go.mod 
module example.com/app

go 1.23.4
root@fedora:~/go_dir/app# go mod tidy
go: found example.com/server in example.com/server v0.0.0-00010101000000-000000000000
go: downloading example.com/utils v0.0.0-00010101000000-000000000000
go: example.com/app imports
        example.com/server imports
        example.com/utils: unrecognized import path "example.com/utils": reading https://example.com/utils?go-get=1: 404 Not Found
root@fedora:~/go_dir/app# go mod edit -replace example.com/utils=../utils/
root@fedora:~/go_dir/app# cat go.mod 
module example.com/app

go 1.23.4

replace example.com/server => ../server/

replace example.com/utils => ../utils/
root@fedora:~/go_dir/app# go mod tidy
go: found example.com/server in example.com/server v0.0.0-00010101000000-000000000000
root@fedora:~/go_dir/app# cat go.mod 
module example.com/app

go 1.23.4

replace example.com/server => ../server/

replace example.com/utils => ../utils/

require example.com/server v0.0.0-00010101000000-000000000000

require example.com/utils v0.0.0-00010101000000-000000000000 // indirect
root@fedora:~/go_dir/app# cat run.go 
package main

import (
        "example.com/server"
)

func main() {
        server.SetupServer("SuperGoServer")
}
root@fedora:~/go_dir/app# go run .
Setting up server with name: "SuperGoServer"
GETTING HOSTNAME...
3...2...1...DONE
HOSTNAME: fedora
IDENTIFYING INTERFACES...
5...4...3...2...1...DONE
NUMBER OF FOUND INTERFACES: 2
INTERFACE: [lo] | IPV4 ADDRESS: [127.0.0.1/8]
INTERFACE: [eth0] | IPV4 ADDRESS: [192.168.211.132/28]
FINISHING SERVER SETUP...
2...1...DONE
All done! Server "SuperGoServer" is ready to be used!
```
