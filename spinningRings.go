package kata

func SpinningRings(innerMax, outerMax int) int {

	//both inner and outer start at 0 (move 0)
	//inner decrements, so it is = innerMax on move 1
	//outer increments, so it is = 1 on move 1

	var moves = 1
	var iVal = innerMax
	var oVal = 1

	for iVal != oVal {
		moves++
		//this is a beautiful way to cycle through a range without the use of if statements
		iVal = (iVal + innerMax) % (innerMax + 1)
		oVal = (oVal + 1) % (outerMax + 1)
	}

	return moves
}
