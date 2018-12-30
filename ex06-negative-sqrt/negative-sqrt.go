package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("can't Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return -1, ErrNegativeSqrt(x)
    }

    EPS := 0.000001
    z0 := 0.0
    z := x/2
    for z - z0 > EPS || z - z0 < -EPS {
        z0 = z
        z -= (z*z - x) / (2*z)
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}


