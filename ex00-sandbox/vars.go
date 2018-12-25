package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt int        = 1<<63 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func typeCasting() {
	a, b := 3, 4
	c := math.Sqrt(float64(a*a + b*b))
	fmt.Printf("Type: %T. Value %v \n", c, c)
}

func constants() {
	const Pi = 3.14
	fmt.Println("Happy", Pi, "Day")
}

func main() {
	fmt.Printf("Type: %T. Value %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T. Value %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T. Value %v \n", z, z)

	typeCasting()
	constants()
}
