package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

type MyFloat float64

// This mean that MyFloat implements Abser interface
// but we don't need to explicitly declare this.
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


// A value of interface type can hold any value that implements those methods.
func interfacesCast() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // a MyFloat implements Abser
    // Under the hood, interface values can be thought of as a tuple of a value
    // and a concrete type
    describe(a)
    fmt.Println(a.Abs())

    a = &v // a *Vertex implements Abser
    describe(a)
    fmt.Println(a.Abs())

    // In the following line, v is a Vertex (not *Vertex)
    // and does NOT implement Abser.
    //a = v
}

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}

// In Go it is common to write methods that gracefully handle being called with
// a nil receiver.
func nilInterface() {
    var i I

    var t *T
    describe(t)

    i = t
    describe(i)
    i.M()

    if t == nil {
        fmt.Println("t is nil")
    }

    // Note that an interface value that holds a nil concrete value is itself
    // non-nil
    if i == nil {
        fmt.Println("i is nil")
    }
}

// An empty interface may hold values of any types.
func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

func typeAssertion() {
    var i interface {} = "string"

    // This statement asserts that the interface value i holds the concrete type T
    // and assigns the underlying T value to the variable t
    s := i.(string)

    fmt.Println(s)

    // To test whether an interface value holds a specific type, a type assertion
    // can return two values: the underlying value and a boolean value that
    // reports whether the assertion succeeded
    s, ok := i.(string)
    fmt.Println(s, ok)

}

func printType(i interface {}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}

func typeSwitches() {
    printType(21)
    printType("hello")
    printType(true)
}

func main() {
    interfacesCast()

    nilInterface()

    typeAssertion()

    typeSwitches()
}
