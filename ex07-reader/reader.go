package main

import (
    "golang.org/x/tour/reader"
    "strings"
)

type MyReader struct{}

func (r *MyReader) Read(b []byte) (n int, err error){
    n = len(b)
    err = nil
    copy(b, strings.Repeat("A", n))
    return
}

func main() {
    reader.Validate(&MyReader{})
}

