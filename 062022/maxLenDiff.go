package kata

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	maxA1 := getMax(a1)
	minA1 := getMin(a1)
	maxA2 := getMax(a2)
	minA2 := getMin(a2)
	if (maxA1 - minA2) > getAbs(minA1-maxA2) {
		return maxA1 - minA2
	} else {
		return getAbs(minA1 - maxA2)
	}

}

func getMax(arr []string) int {
	max := 0
	for _, i := range arr {
		if len(i) > max {
			max = len(i)
		}
	}
	return max
}

func getMin(arr []string) int {
	min := 0
	for _, k := range arr {
		if len(k) < min || min == 0 {
			min = len(k)
		}
	}
	return min
}

func getAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
