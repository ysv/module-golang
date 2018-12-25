package main

import "fmt"

func rangeIterating(){
    arr := []int{1, 2, 4, 8, 16, 32, 64}
    for i, pow := range arr{
        fmt.Printf("2**%d = %d", i, pow)
    }
}

func main(){
    rangeIterating()
}
