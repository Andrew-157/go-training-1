/*
Second implementation of package that implements Echo function,
which is a very primitive implementation of Unix-like echo command
*/
package echo2

import (
	"fmt"
	"os"
)

// example.com/echo2: Prints command-line arguments with space as a separator
func Echo() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
