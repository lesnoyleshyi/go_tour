package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go doesn't support classes, but it support methods
// A method is a function with a special argument called receiver
// In this example method is called Abs()
// (v Vertex) is so called receiver. We will call Abs() method on it: v.Abs()
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// As methods in go are just functions, Abs() can be rewritten this way:
func also_abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// We can define methods with non-struct types, but those types must be specified in
// the same package as method. Therefore we can't use regular types, but we can define our own
type Myint int64

func (num Myint) Summator() int64 {
	return int64(num + 1)
}

func main() {
	my_struct := Vertex{3, 4}
	var number Myint = 1488
	fmt.Println(my_struct.Abs())
	fmt.Println(also_abs(my_struct))
	fmt.Println(number.Summator())
}
