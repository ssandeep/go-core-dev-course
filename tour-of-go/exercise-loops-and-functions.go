package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	count := 0
	for math.Abs(z*z-x) > 0.00001 {
		count++
		z -= (z*z - x) / (2 * z)
	}
	fmt.Println("No of iterations:", count)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
