package main

import "fmt"

func main() {
	var sum int = 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 100 {
		sum += sum
	}
	fmt.Println("New sum", sum)
	fmt.Println("There are two forms: 'for ; sum < 100; ' and 'for sum < 100'")

	//for {} - it's infinite loop
}
