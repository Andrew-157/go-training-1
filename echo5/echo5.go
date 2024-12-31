/*
Fifth implementation of package that implements Echo function
*/
package echo5

import (
	"fmt"
	"os"
)

// example.com/echo5: Prints the index and value of each of its arguments, one per line
func Echo() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("%d - %s\n", index, arg)
	}
}
