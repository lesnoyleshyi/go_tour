package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var logical bool = true
	var str string = "He-he"
	var usualint_min int = -2147483648
	var usualint_max int = 2147483647
	var (
		int8_min  int8   = -128
		int8_max  int8   = 127
		int16_min int16  = -32768
		int16_max int16  = 32767
		int32_min int32  = -2147483648
		int32_max int32  = 2147483647
		int64_min int64  = -9223372036854775808
		int64_max int64  = 9223372036854775807
		maxint    uint64 = 1<<64 - 1
	)

	var lilfloat float32 = 100.5
	var bigfloat float64 = 10.5

	var y complex128 = cmplx.Sqrt(-10 + 20i)
	var z complex64

	fmt.Println("Regular int (int, uint, uintptr) range depends on system: it's int32 for 32-bit systems"+
		"and int64 for 64-bit systems", int8_min, int8_max, "\n", int16_min, int16_max, "\n",
		int32_min, int32_max, "\n", int64_min, int64_max)
	fmt.Println(str, usualint_min, usualint_max)
	fmt.Println(logical)
	fmt.Println("Byte is alias for int8, ", int8_min, int8_max,
		"\nwhereas rune is an alias for int32", int32_min, int32_max)
	fmt.Printf("Float32 occupies 32 bit of memory: %v, %T\n", lilfloat, lilfloat)
	fmt.Printf("Float64 occupies 32 bit of memory: %v, %T\n", bigfloat, bigfloat)
	fmt.Printf("The biggest number has type %T and vallue %v\n", maxint, maxint)
	fmt.Printf("There is also %T and %T types, examples: %v %v\n", z, y, z, y)

}
