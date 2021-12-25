package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	if s == nil {
		fmt.Println("It is nill")
	}
}

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 10, 20, 30, 40, 50)
	printSlice(s)

	for i, v := range s {
		fmt.Printf("Num: %d, val: %d\n", i, v)
	}

	for i, _ := range s {
		fmt.Printf("Num: %d\n", i)
	}

	for _, v := range s {
		fmt.Printf("Val: %d\n", v)
	}

	for i := range s {
		fmt.Println(i)
	}
}
