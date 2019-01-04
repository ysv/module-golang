package main

import (
    "fmt"
    "math/rand"
    "time"
)

func goSum(s []int, res chan int) {
    sum := 0
    for _, n := range s {
        sum += n
    }
    res <- sum
}


func concurrentSum() {
    rand.Seed(time.Now().UnixNano())
    size := rand.Int() % 15
    numbers := rand.Perm(size)

    //By default, sends and receives block until the other side is ready.
    // This allows goroutines to synchronize without explicit locks
    // or condition variables.
    res := make(chan int)
    go goSum(numbers[size/2:], res)
    go goSum(numbers[:size/2], res)

    sum1, sum2 := <-res, <-res

    fmt.Println(size)

    fmt.Println(numbers)
    fmt.Print(sum1, sum2, sum1 + sum2)
}

func bufferedChannels() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func closeChannels() {
    c := make(chan int, 5)
    go fibonacci(10, c)
    for i := range c {
        fmt.Println(i)
    }
}

func put(n int, c chan int) {
    for i := 0; i < n; i++ {
        c <- rand.Intn(100)
    }
    close(c)
}

func closeChannels2() {
    c := make(chan int)
    expect := 5
    go put(expect, c)
    for i := 0; i < expect + 1; i++ {
        num, ok := <- c
        fmt.Println(num, ok)
        if ok != true {
            fmt.Print("not ok !!! Break")
            break
        }
    }
}

func fibonacci2(c, quit chan int) {
    x, y := 0, 1
    for {
    // The select statement lets a goroutine wait on multiple communication
    // operations. A select blocks until one of its cases can run, then it
    // executes that case. It chooses one at random if multiple are ready
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func selectStatement() {
    c := make(chan int, 2)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci2(c, quit)
}

func selectStatement2() {
    tick := time.Tick(1000 * time.Millisecond)
    boom := time.After(5000 * time.Millisecond)
    for {
        select {
        case t := <-tick:
            fmt.Printf("tick. %v", t)
        case t := <-boom:
            fmt.Printf("BOOM! %v", t)
            return
        default:
            fmt.Println("    .")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    concurrentSum()

    bufferedChannels()

    closeChannels()

    closeChannels2()

    selectStatement()

    selectStatement2()
}

