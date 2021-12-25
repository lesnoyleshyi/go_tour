package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	//Point.X = 12
	//Point.Y = 21
	//fmt.Println(Point.X, Point.Y)
	fmt.Println(Point{12, 21})

	a := Point{10, 20}
	a.X = 5
	fmt.Println(a.Y)

	pointer := &a
	pointer.X = 15
	fmt.Println(a)
	//So we refer to structs fields by dot whether it
	// pointer to structure or regular structure

}
