package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Error() string {
	return fmt.Sprintf("Something went wrong with MyReader %v", r)
}

func (r MyReader) Read(b []byte) (int, error) {
	var ret int
	for i, _ := range b {
		ret = i
		b[i] = 'A'
	}
	return ret, r
}

func main() {
	reader.Validate(MyReader{})
}
