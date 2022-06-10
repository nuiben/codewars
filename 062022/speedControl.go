package kata

func Gps(s int, x []float64) int {

	avg := 0.0
	prev := 0.0
	maxVal := 0

	for _, i := range x {
		avg = (3600 * (i - float64(prev))) / float64(s)
		if int(avg) > maxVal {
			maxVal = int(avg)
		}
		prev = i
	}
	return maxVal
}
