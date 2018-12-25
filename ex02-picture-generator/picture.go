package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (arr [][]uint8) {
    arr = make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        arr[i] = make([]uint8, dx)
        for j:= 0; j < dx; j++ {
            arr[i][j] = uint8(i * j)
        }
    }
    return
}

func main() {
    pic.Show(Pic)
}

