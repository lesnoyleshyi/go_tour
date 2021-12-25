package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("world", time.Now())
	fmt.Print("Hello, ", time.Now(), "\n")
}
