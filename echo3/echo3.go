/*
Third implementation of package that implements Echo function,
which is a very primitive implementation of Unix-like echo command
*/
package echo3

import (
	"fmt"
	"os"
	"strings"
)

// example.com/echo3: Prints command-line arguments with space as a separator
func Echo() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
