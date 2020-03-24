package iterator

func Repeat(input string, times int) (result string) {
	for i := 0; i < times; i++ {
		result += input
	}
	return
}
