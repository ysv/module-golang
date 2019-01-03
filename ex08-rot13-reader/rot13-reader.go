package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {
    n, err := rot13.r.Read(b)
    for i, c := range b{
        if c >= 'a' && c <= 'z' {
            c = (c + 13) % 'z'
            if c < 'a' {
                c += 'a' - 1
            }
        }
        if c >= 'A' && c <= 'Z' {
            c = (c + 13) % 'Z'
            if c < 'A' {
                c += 'A' - 1
            }
        }
        b[i] = c
    }
    return n, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

