package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	S := scanner.Text()
	_ = S // to avoid unused error

	explode := strings.Split(S, "")

	var output string = "("

	for idx, character := range explode {

		if idx == 3 {
			output += ") "
		}
		if idx == 6 {
			output += "-"
		}

		output += character
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(output) // Write answer to stdout
}
