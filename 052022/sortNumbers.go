package kata

func SortNumbers(numbers []int) []int {
	if numbers == nil {
		return nil
	}

	for i := range numbers {
		for j := range numbers {
			if numbers[i] < numbers[j] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}

	return numbers
}
