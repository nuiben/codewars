package kata

import (
	"fmt"
	"sort"
)

func SumOfDivided(lst []int) string {
	var sums = make(map[int]int)

	for _, n := range lst {
		positiveN := n
		if n < 0 {
			positiveN = n * -1
		}
		for i := 2; i*i <= positiveN; i++ {
			if positiveN%i == 0 {
				sums[i] += n
				for positiveN%i == 0 {
					positiveN /= i
				}
			}
		}
		if positiveN > 1 {
			sums[positiveN] += n
		}
	}

	keys := []int{}
	for key := range sums {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	result := ""
	for _, key := range keys {
		result += fmt.Sprintf("(%d %d)", key, sums[key])
	}
	return result
}
