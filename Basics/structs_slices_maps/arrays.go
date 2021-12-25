package main

import "fmt"

func make_diagonal(len int, sq_arr [][]string) {
	for i := 0; i < len; len++ {
		for j := 0; j < len; len++ {
			sq_arr[i][j] = "X"
		}
	}
}

func main() {
	var a [10]string
	a[0] = "Hello"
	a[1] = ", "
	a[2] = "World"
	a[3] = "!"

	b := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < 7; i++ {
		fmt.Println(a[i])
	}

	fmt.Println(a)
	fmt.Println(b)
	//Arrays are fix-sized, when slices are dynamically-sized
	c := b[1:3]

	fmt.Println(c)
	c[0] = 10
	fmt.Println(b)
	//We can change original array through slice

	d := []int{2, 4, 6, 8, 10, 12, 14}
	fmt.Println(d)
	//When we don't provide size in [] we create both:
	//slice explicitly and array implicitly

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	t := s[0:6]
	t1 := s[0:]
	t2 := s[:4]
	t3 := s[:]

	fmt.Println(t, t1, t2, t3)
	fmt.Println(len(t2), cap(t2))

	var m []int
	fmt.Println(m, len(m), cap(m))
	if m == nil {
		fmt.Println("nil")
	}
	new_slice := make([]int, 5, 10)
	fmt.Printf("%T %v\n", new_slice, new_slice)
	new_slice = append(new_slice, 33)
	fmt.Printf("%T %v %d\n", new_slice, new_slice, cap(new_slice))

	real_board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	slice_of_board := real_board[:]
	// fmt.Println(make_diagonal(3, slice_of_board))
	// fmt.Println(make_diagonal(3, real_board))
	fmt.Println(slice_of_board)
}
