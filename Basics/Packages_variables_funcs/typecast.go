package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 5, 10
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	w := uint(f)

	fmt.Printf("To cast something we have to write CAST_TYPE(expression/variable):\n"+
		"This is f with type %T: %v\n", f, f)
	fmt.Printf("And this is z with type %T: %v\n", z, z)
	fmt.Printf("w and z variables are the same: %v, %v\n", w, z)
	fmt.Println("Unlike C, we must cast one type to another explicitly\n" +
		"When assign to variable value of different type")
}
