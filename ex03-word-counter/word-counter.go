package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCounter(s string) map[string]int {
    counter := make(map[string]int)

    for _, w := range strings.Fields(s){
        counter[w] += 1
    }
    return counter
}

func main() {
    wc.Test(WordCounter)
}

