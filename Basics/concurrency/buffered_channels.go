package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	// ch <- 3 This will block channel, cause its size is 2
	// and noone can "free" this channel
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) This will cause an error - channel doesn't have anything
	ch <- 3
	fmt.Println(<-ch)
}
