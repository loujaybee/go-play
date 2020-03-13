package main

import (
    "fmt"
    "flag"
)

func main() {

    var flagvar int
	flag.IntVar(&flagvar, "number", 0, "Number for fibonnaci")
    flag.Parse()

    result := fibonnaci(flagvar)

    fmt.Println(result)
}

// 0 = 0
// 1 = 1
// 2 = 1
// 3 = 2
// 4 = 3
// 5 = 5
// 6 = 8
