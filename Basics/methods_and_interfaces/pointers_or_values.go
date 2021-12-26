package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// It's method with pointer receiver
func (v *Vertex) p_scaler(scale int) {
	v.X *= float64(scale)
	v.Y *= float64(scale)
}

// It's method with value receiver
func (v Vertex) v_scaler(scale int) {
	v.X *= float64(scale)
	v.Y *= float64(scale)
}

// It's function that receive pointer to variable
func f_p_scaler(v *Vertex, scale int) {
	v.X *= float64(scale)
	v.Y *= float64(scale)
}

// It's function that receive variable
func f_v_scaler(v Vertex, scale int) Vertex {
	res := Vertex{v.X * float64(scale), v.Y * float64(scale)}
	return res
}

func main() {
	v := Vertex{3, 4}
	p := &v

	v.p_scaler(2)
	fmt.Println(v)
	v.v_scaler(2)
	fmt.Println(v)
	p.p_scaler(2)
	fmt.Println(v)
	p.v_scaler(2)
	fmt.Println(v)
	// It'll all work: methods can receive either pointer or value
	// No matter whether they has pointer receiver or value receiver
	// But when we use value receiver, method can't change value of variable
	// from which method was called

	// f_p_scaler(v, 10) will not work as this function can only receive pointers
	f_p_scaler(p, 10)
	fmt.Println(v)

	// f_v_scaler(p, 10) will not work as this function can only receive values
	v = f_v_scaler(v, 10)
	fmt.Println(v)

	// Here are two reasons to use pointer receiver from Go tour:
	// The first is so that the method can modify the value that its receiver points to.
	// The second is to avoid copying the value on each method call.
	//  This can be more efficient if the receiver is a large struct, for example.
}
