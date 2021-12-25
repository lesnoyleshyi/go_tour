package main

import (
	"fmt"
	"math"
)

func my_sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return 0
}

func my_sqrt2(x float64) float64 {
	rez := 1.0
	for i := 1; math.Abs(math.Sqrt(x)-rez) > 0.0000000001; i++ {
		rez -= (rez*rez - x) / (2 * rez)
		fmt.Println(rez, i)
	}
	return rez
}

func main() {
	fmt.Println(my_sqrt(2))
	fmt.Println("And the second func: ")
	fmt.Println(my_sqrt2(2))
}
