package main

import (
	"flag"
	"fmt"
)

func main() {

	team := flag.String("team", "No Team Specified", "The team name")

	flag.Parse()

	result := businesslogic(*team)

	fmt.Println(result)
}

// 0 = 0
// 1 = 1
// 2 = 1
// 3 = 2
// 4 = 3
// 5 = 5
// 6 = 8
