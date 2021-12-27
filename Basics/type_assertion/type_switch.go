package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Value is of %T type: v*v=%v\n", v, v*v)
	case string:
		fmt.Printf("Value is of %T type: len(s)=%v\n", v, len(v))
	default:
		fmt.Printf("Value if of niether string no int. It's %T\n", v)
	}
}

func main() {
	do(10)
	do("Hello, world!")
	do(3.1415)
}
