/*
Fourth implementation of package that implements Echo function,
which is a very primitive implementation of Unix-like echo command
and also differs from normal echo a bit
*/
package echo4

import (
	"fmt"
	"os"
	"strings"
)

// example.com/echo4: Prints name of the executable that runs the code and then
// prints command-line arguments with space as a separator
func Echo() {
	fmt.Println("File that executes Echo:", os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
