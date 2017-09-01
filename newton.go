/*
Implementation of square root function using newtons method of apporximating point z and then repeating: z - (z*z - x) / (2 * z)
*/

package main

import (
	"fmt"
	"math"
)

var p = 0.0

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(p - z) > 0.0000000000001 { 
		z = newton(z, x)
		p = newton(z, x)
	}
	return z
}

func newton(z, x float64) float64 {
	return z - ( ((z*z) - x) / (2*z) )
}

func main() {
	fmt.Println(Sqrt(4))
}
