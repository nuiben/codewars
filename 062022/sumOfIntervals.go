package kata

func SumOfIntervals(intervals [][2]int) int {
	dict := make(map[int]int)
	for _, i := range intervals {
		for j := i[0]; j < i[1]; j++ {
			dict[j] += 1
		}
	}
	return len(dict)

}
