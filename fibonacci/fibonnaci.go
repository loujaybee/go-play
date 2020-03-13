package main

func fibonnaci(flagvar int) int {
	first := 1
	second := 0
	curr := 0

	for i := 0; i < flagvar; i++ {
		// fmt.Println(
		//     "| passed: ", flagvar,
		//     "| first: ", first,
		//     "| second: ", second,
		//     "| i: ", i,
		//     "|",
		// )
		curr = first
		first = first + second
		second = curr
	}

	return curr
}
