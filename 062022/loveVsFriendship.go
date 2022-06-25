package kata

func WordsToMarks(s string) int {
	var result int
	for _, i := range s {
		result += int((i - 96))
	}
	return result
}
