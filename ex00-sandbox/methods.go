package main

import (
    "fmt"
    "math"
)

// A method is a function with a special receiver argument
type Vector struct {
    X, Y float64
}

func (v Vector) Length() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Methods with pointer receivers can modify the value to which the receiver points
func (v *Vector) Scale(f float64) {
    v.X *= f
    v.Y *= f
}

type MyFloat float64

// You cannot declare a method with a receiver whose type
// is defined in another package
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func vectorScale() {
    v := Vector{3, 4}
    v.Scale(10)
    fmt.Println(v.Length())
}

func main() {
    v := Vector{3, 4}
    fmt.Println(v.Length())

    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())

    vectorScale()
}
