package kata

func PositiveSum(numbers []int) int {
	var sum int = 0
	for i := range numbers {
		if numbers[i] > 0 {
			sum = sum + numbers[i]
		}
	}
	return sum
}
