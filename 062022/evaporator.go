package kata

func Evaporator(content float64, evapPerDay int, threshold int) int {

	day := 0
	newThresh := content * (float64(threshold) / 100)
	for i := 0; content >= newThresh; i++ {
		content -= content * (float64(evapPerDay) / 100)
		day++
	}

	return day
}
