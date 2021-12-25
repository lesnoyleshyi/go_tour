package main

import "fmt"

func main() {
	i, j := 42, 142

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)

	var r *int

	r = &j

	fmt.Println(*r)
	//pointer arifmetic is not supported by Go

}
