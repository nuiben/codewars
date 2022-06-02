package kata

func TwoSum(numbers []int, target int) [2]int {
	var answer [2]int
	for x, _ := range numbers {
		for y, _ := range numbers {
			if numbers[x]+numbers[y] == target && x != y {
				answer[0] = x
				answer[1] = y
				return answer
			}
		}
	}
	return answer
}
