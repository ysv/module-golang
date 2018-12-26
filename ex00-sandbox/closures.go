package main

import "fmt"

func closure() func(int) int {
    sum := 0
    return func(x int) int {
        if x % 2 == 0{
            x = -x
        }
        sum += x
        return sum
    }
}
func main() {
    odd, even := closure(), closure()
    for i := 0; i < 10; i++ {
        if i % 2 == 1 {
            odd(i)
        } else {
            even(i)
        }
    }
    fmt.Printf("odd = %d\n", odd(0))
    fmt.Printf("even = %d\n", even(0))


}
