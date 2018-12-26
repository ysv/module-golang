package main

import "fmt"

func fibonacci() func() int {
    a1, a2 := -1, 1
    return func() int {
        a1, a2 = a2, a1 + a2
        return a2
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

