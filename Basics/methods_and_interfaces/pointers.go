package main

import "fmt"

type Vertex struct {
	X, Y int64
}

// Methods with pointer receivers are widespread
// Such methods can change value of "object" from which this method was called
func (v *Vertex) Scaler(scale int) {
	v.X *= int64(scale)
	v.Y *= int64(scale)
}

// Nevertheless, methods are functions, thus Scaler can be rewritten:
func f_scaler(v *Vertex, scale int) {
	v.X *= int64(scale)
	v.Y *= int64(scale)
}

func main() {
	lol := Vertex{3, 4}

	fmt.Println(lol)
	lol.Scaler(10)
	fmt.Println(lol)

	lol.Scaler(2)
	fmt.Println(lol)
	f_scaler(&lol, 5)
	fmt.Println(lol)
}
