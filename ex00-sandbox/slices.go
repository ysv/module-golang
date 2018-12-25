package main

import (
    "fmt"
)


func modifySlices(){
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names)

    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)

    b[0] = "XXX"
    fmt.Println(a, b)
    fmt.Println(names)
    fmt.Printf("%T", a)
}

func lenAndCapacity(){
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    // Slice the slice to give it zero length.
    s = s[:0]
    printSlice(s)

    // Extend its length.
    s = s[:4]
    printSlice(s)

    // Drop its first two values.
    s = s[2:]
    printSlice(s)

    //The capacity of a slice is the number of elements in the underlying array,
    //counting from the first element in the slice.
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilSlice(){
    var s []int
    fmt.Println(s, len(s), cap(s))
    if s == nil{
        fmt.Println("s == nil")
    }
}

func makeSlice(){
    // 3 - size, 5 - capacity.
    s := make([]int, 3, 5)
    printSlice(s)
}

func appendSlice(){
    s := []int{-1, 0, 1}
    s = append(s, 2, 3, 4)
    printSlice(s)
}

func sliceGotcha(){
    //var digitRegexp = regexp.MustCompile("[0-9]+")
    //
    //func FindDigits(filename string) []byte {
    //    b, _ := ioutil.ReadFile(filename)
    //    return digitRegexp.Find(b)
    //}
    // This code behaves as advertised, but the returned []byte
    // points into an array containing the entire file
    // Since the slice references the original array,
    // as long as the slice is kept around the garbage collector can't release
    // the array;
}

func main() {
    modifySlices()
    //var a [10]int
    //these slice expressions are equivalent
    //a[0:10]
    //a[:10]
    //a[0:]
    //a[:]

    lenAndCapacity()

    nilSlice()

    makeSlice()

    appendSlice()
}
