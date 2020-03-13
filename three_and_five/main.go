package main

import (
	"flag"
	"fmt"
)

func main() {

	var flagvar int
	flag.IntVar(&flagvar, "number", 0, "Input argument")
	flag.Parse()

	result := threeAndFive(flagvar)

	fmt.Println(result)
}
