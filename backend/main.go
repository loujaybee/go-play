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
