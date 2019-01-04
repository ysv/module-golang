package main

import (
    "fmt"
    "time"
)

func say(word string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(word)
    }
}

// Goroutines run in the same address space,
// so access to shared memory must be synchronized
func main() {
    go say("world")
    say("hello")
}
