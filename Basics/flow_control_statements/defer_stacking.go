package main

import "fmt"

func main() {
	fmt.Println("He-he")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done execution")
}
