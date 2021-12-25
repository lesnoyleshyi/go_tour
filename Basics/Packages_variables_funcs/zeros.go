package main

import "fmt"

func main() {
	var (
		i uintptr
		f float64
		b bool
		s string
	)
	fmt.Println(i, f, b, "->", s, "<-")
}
