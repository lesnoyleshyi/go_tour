package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Print("When is Saturday?\n")
	today := time.Now().Weekday()
	fmt.Println("Today's value: ", today)
	fmt.Println("time.Saturday value: ", time.Saturday)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
	// case stops when first switch goes true

	switch {
	case time.Now().Hour() < 12:
		fmt.Println("Good morning!")
	case time.Now().Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
	// There is not condition in switch == switch true
}
