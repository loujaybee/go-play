package main

func threeAndFive(flagvar int) int {

	sum := 0

	for i := 0; i < flagvar; i++ {
		if (i % 3 == 0 || i % 5 == 0) {
			sum += i
		}
	}

	return sum
}
