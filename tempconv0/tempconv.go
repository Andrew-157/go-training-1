// Package tempconv0 performs Celsius and Fahrenheit temperature computations.
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // 100
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // 180
	// fmt.Printf("%g\n", boilingF-FreezingC) // (mismatched types Fahrenheit and Celsius)

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	// fmt.Println(c == f) // (mismatched types Celsius and Fahrenheit)
	fmt.Println(c == Celsius(f)) // true
}
