package kata

func GetSize(w, h, d int) [2]int {
	area := w*h*2 + d*h*2 + d*w*2
	volume := w * h * d
	ans := [2]int{area, volume}
	return ans
}
