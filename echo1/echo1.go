/*
First implementation of package that implements Echo function,
which is a very primitive implementation of Unix-like echo command
*/
package echo1

import (
	"fmt"
	"os"
)

// example.com/echo1: Prints command-line arguments with space as a separator
func Echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
