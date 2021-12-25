package main

import (
	"fmt"
	// "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	fmt.Println("It's what we get after initialization:", arr)
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			arr[i] = append(arr[i], uint8((j+i)/2))
		}
	}
	return arr
}

func main() {
	// pic.Show(Pic)
	fmt.Println(Pic(5, 6))
}
