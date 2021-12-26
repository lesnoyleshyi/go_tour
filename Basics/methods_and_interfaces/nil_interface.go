package main

import "fmt"

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	describe(i)
	// The next line will cause an error
	i.M()
	// Calling a method on a nil interface is a run-time error because there is no type
	// inside the interface tuple to indicate which concrete method to call.
}
