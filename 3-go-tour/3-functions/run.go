package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("The sum of 34 and 68 is", add(34, 68))
	fmt.Println("The difference of 34 and 68 is", subtract(34, 68))
	x, y := swap("world", "hello")
	fmt.Println(x, y)
	fmt.Println(split(17))
}
