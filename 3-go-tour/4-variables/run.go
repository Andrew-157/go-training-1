package main

import "fmt"

var c, python, java bool
var x, y int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java) // Output: 0 false false false
	var cpp, ruby, rust = true, false, "no!"
	fmt.Println(x, y, cpp, ruby, rust) // Output: 1 2 true false no!

	a := 45
	v, d, f := true, false, "no!"
	fmt.Println(a, v, d, f) // Output: 45 true false no!
}
