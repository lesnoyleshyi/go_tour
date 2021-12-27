package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(symb byte) (coded_symb byte) {
	if (symb >= 'a' && symb <= 'm') || (symb >= 'A' && symb <= 'M') {
		coded_symb = symb + 13
	} else if (symb > 'm' && symb <= 'z') || (symb > 'M' && symb <= 'Z') {
		coded_symb = symb - 13
	} else {
		coded_symb = symb
	}
	return
}

func (rot *rot13Reader) Read(p []byte) (int, error) {
	ret, err := rot.r.Read(p)
	if err != nil {
		return 0, err
	}
	for i, v := range p {
		p[i] = rot13(v)
	}
	return ret, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println("")
}
