package kata

import "fmt"

func SumOfIntervals(intervals [][2]int) int {
	fmt.Println("input: ", intervals)
	fmt.Println("len: ", len(intervals))
	sum := 0
	next := 0
	//   mrgIntervals := intervals

	for k, i := range intervals {
		if k+1 < len(intervals) {
			next = k + 1
		}
		fmt.Println("next: ", next)
		fmt.Println("k: ", k, "   i: ", i)
		//     fmt.Println("intervals[next][0]: ",intervals[next][0])
		//     fmt.Println("i[0]: ",i[0])
		if k != next && intervals[next][0] <= i[1] {
			if intervals[next][0] < i[0] {
				sum += i[1] - i[0]
				sum += intervals[next][1] - intervals[0][1]
				sum -= intervals[next][1] - intervals[next][0]
				fmt.Println("the next interval doesn't overlap the current, sum: ", sum)
			} else {
				sum += intervals[next][1] - i[0]
				fmt.Println("overlapping count: ", sum)
				break
			}
		} else {
			sum += i[1] - i[0]
			fmt.Println("normal count: ", sum)
		}
	}
	return sum
}
