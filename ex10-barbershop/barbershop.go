package main

import (
    "fmt"
    "math/rand"
    "time"
)


type Visitor struct {
    fName string
    sName string
}

type Barber struct {
    isSleeping bool
}

type Queue struct {
    waitQueue chan Visitor
}

type BarberShop struct {
    queue *Queue
    barber *Barber
}


func NewBarber() *Barber {
    return &Barber{false}
}

func NewQueue() *Queue {
    return &Queue{make(chan Visitor, 3)}
}

func NewBarberShop() *BarberShop {
    return &BarberShop{NewQueue(), NewBarber()}
}

func (bs *BarberShop) startToWork() {
    for {
        select {
        case visitor := <-bs.queue.waitQueue:
            fmt.Println("Welcome ", visitor)
            time.Sleep(500 * time.Millisecond)
        default:
            bs.barber.isSleeping = true
            fmt.Println("Lets have a sleep")
            time.Sleep(800 * time.Millisecond)
        }
    }
}

func (v Visitor) visit(bs *BarberShop) {
    bs.queue.waitQueue <- v
    fmt.Println(v, " is on queue")
    time.Sleep(200 * time.Millisecond)
}

func main() {
    names := []string{"A", "B", "C", "D"}
    bs := NewBarberShop()
    go bs.startToWork()
    for {
       v := Visitor{fName: names[rand.Intn(4)]}
       v.visit(bs)
    }
}
