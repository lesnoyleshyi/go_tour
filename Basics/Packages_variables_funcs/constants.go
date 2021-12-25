package main

import "fmt"

const Pi = 3.1415

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const World string = " "
	const (
		Truth  bool = true
		symbol int8 = 'x'
		num         = 1000
		Big         = 1 << 100
		Small       = Big >> 99
	)

	fmt.Println(Pi, World, Truth, symbol, num)
	fmt.Println("We can't define variable with := notation")

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	//fmt.Println(needInt(Big)) - this throw an error
	fmt.Println(needFloat(Big))
}
