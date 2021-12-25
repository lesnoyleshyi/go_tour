package main

import "fmt"

var c, python, java bool
var golang, cpp bool = true, true

func main() {
	var i int
	var j float32 = 14.88
	var k, word = 13.77, "Good"
	m, another_word := 100, "Bad"
	fmt.Println("Variables have standard values without explicit initialization: ",
		i, c, python, java)
	fmt.Println("They can also be in different scopes:\n"+
		"In the level of package: golang", golang,
		"c++", cpp,
		"\nOr in the scope of function: ", j)
	fmt.Println("If we use initialisers, we shouldn't provide type of variable "+
		"explicitly : ", word, k)
	fmt.Println("Short noration := is available only inside function: ", m, another_word)
}
