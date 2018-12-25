package main

import "fmt"

type MapVertex struct {
    Lat, Long float64
}

var m = map[string]MapVertex{
    "Bell Labs": MapVertex{
        40.68433, -74.39967,
    },
    "Google": { 37.42202, -122.08408 },
}

func mapLiterals() {
    fmt.Println(m)
}

func mutatingMaps() {
    mp := make(map[string]int)

    mp["Answer"] = 42
    fmt.Println("The value:", mp["Answer"])

    delete(mp, "Answer")
    fmt.Println("The value:", mp["Answer"])

    // Test value presence.
    v, ok := mp["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}

func main() {
    mapLiterals()

    mutatingMaps()
}
