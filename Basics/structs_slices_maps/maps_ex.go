package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var m map[string]int
	m = make(map[string]int)
	for _, word := range strings.Fields(s) {
		if _, ok := m[word]; ok == true {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}
	return m
}

func main() {
	string := "some string with some words"
	fmt.Println(WordCount(string))
}
