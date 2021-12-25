package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y, z int) int {
	return x + y + z
}

func sort(a, b, c int) (int, int, int) {
	if a > b {
		if a > c {
			if c > b {
				return a, c, b
			} else {
				return a, b, c
			}
		} else {
			return c, a, b
		}
	} else {
		if c > b {
			return c, b, a
		} else {
			if c > a {
				return b, c, a
			} else {
				return b, a, c
			}
		}
	}
}

func rev_sort(a, b, c int) (max, mid, min int) {
	if a > b {
		if a > c {
			if c > b {
				max = a
				mid = c
				min = b
			} else {
				max = a
				mid = b
				min = c
			}
		} else {
			max = c
			mid = a
			min = b
		}
	} else {
		if c > b {
			max = c
			mid = b
			min = a
		} else {
			if c > a {
				max = b
				mid = c
				min = a
			} else {
				max = b
				mid = a
				min = c
			}
		}
	}
	return min, mid, max
}

func main() {
	fmt.Println("Result of add is", add(20, 30))
	fmt.Println("Result of add2 is", add2(20, 30, 40))
	first, second, third := sort(25, 75, 10)
	fmt.Println("Result of ugly sort is ", first, second, third)
	first, second, third = rev_sort(25, 75, 10)
	fmt.Println("Result of once more uly sort is ", first, second, third)
}
