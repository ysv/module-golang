package main

import (
    "fmt"
    "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }
    Walk(t.Left, ch)
    ch <- t.Value
    Walk(t.Right, ch)
}

func initChannels(n int) []chan int {
    channels := make([]chan int, n)
    for i := 0; i < n; i++ {
        channels[i] = make(chan int)
    }
    return channels
}

func startWalkRoutines(trees []*tree.Tree, channels []chan int) {
    for i, tr := range trees {
        // We need to pass channel as a parameter into closure.
        // Otherwise this will cause:
        // fatal error: all goroutines are asleep - deadlock!
        // Because it is consider no go routine for channel.
        go func(ch chan int) {
           Walk(tr, ch)
           // Close channel if we finished tree walk.
           close(ch)
        }(channels[i])
    }
}

func assertSameValues(channels []chan int) bool {
    same := true
    finish := false

    for {
        pValue, pOk := <-channels[0]

        for _, ch := range channels[1:] {
            value, ok := <-ch
            // All consider to be the same if:
            // 1. There were no different values before.
            // 2. There are equal values in channel.
            // 3. Tree nodes finishes at the same time  
            same = same && value == pValue && ok == pOk
            finish = finish || !ok
        }
        if finish || !same {
            return same
        }
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(trees []*tree.Tree) bool {
    channels := initChannels(len(trees))
    startWalkRoutines(trees, channels)

    return assertSameValues(channels)
}

func main() {
    trees := []*tree.Tree{tree.New(4), tree.New(4), tree.New(4)}
    fmt.Println(Same(trees))

    trees[0].Value = -1
    fmt.Println(Same(trees))

    trees2 := []*tree.Tree{tree.New(2), tree.New(3), tree.New(8)}
    fmt.Println(Same(trees2))

}
