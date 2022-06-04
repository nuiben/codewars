package kata

func Snail(snaipMap [][]int) []int {

	var b_wid = len(snaipMap[0])
	var b_len = len(snaipMap[0])

	var result = []int{}
	var dir string = "right"
	var xmin int = 0
	var xmax int = b_wid - 1
	var ymin int = 0
	var ymax int = b_len - 1
	var x int = xmin
	var y int = ymin
	var done bool = false

	for !done {
		switch dir {
		case "right":
			x = xmin
			y = ymin
			for y := ymin; y <= ymax; y++ {
				result = append(result, snaipMap[x][y])
			}
			dir = "down"
			xmin++
			break
		case "left":
			x = xmax
			y = ymax
			for y := ymax; y >= ymin; y-- {
				result = append(result, snaipMap[x][y])
			}
			dir = "up"
			xmax--
			break
		case "down":
			x = xmin
			y = ymax
			for x := xmin; x <= xmax; x++ {
				result = append(result, snaipMap[x][y])
			}
			dir = "left"
			ymax--
			break
		case "up":
			x = xmax
			y = ymin
			for x := xmax; x >= xmin; x-- {
				result = append(result, snaipMap[x][y])
			}
			dir = "right"
			ymin++
			break
		}

		if xmin > xmax || ymin > ymax {
			done = true
		}
	}
	return result
}
