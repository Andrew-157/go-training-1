package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Type: %T, Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)

	// Zero Values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s) // Output: 0 0 false ""

	// Type Conversions
	var x, y int = 3, 4
	var e float64 = math.Sqrt(float64(x*x + y*y))
	var z uint64 = uint64(e)
	fmt.Println(x, y, z)

	v := 42
	fmt.Printf("Type of v is %T\n", v) // Type of v is int
	l := 3.142
	fmt.Printf("Type of l is %T\n", l) // Type of l is float64
	r := 0.867 + 0.5i
	fmt.Printf("Type of r is %T\n", r) // Type of r is complex128

	// Constants
	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big))
}
