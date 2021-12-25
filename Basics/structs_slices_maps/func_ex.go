package main

import "fmt"

func fibonacci() func() int {
	var prev int = 0
	var curr int = 1
	return func() int {
		sum := curr + prev
		prev = curr
		curr = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
