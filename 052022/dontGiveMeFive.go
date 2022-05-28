package kata

import (
	"math"
)

func DontGiveMeFive(start int, end int) int {

	//Solution is trickier to solve if digits other than last are '5'
	//ex: 50, 500, 2050, 30500
	//to account for this, we need to get the length and divide the number by 10 incrementally
	//until we've checked all digits, at each increment, we perform the Abs(%10) to see if it contains a '5'.
	//Optimization: We don't need to continually check for length, once we have the length, we only need to update length
	//about the exponents of 10 {10^1, 10^2, 10^3, etc.}

	count := 0

	for i := start; i <= end; i++ {

		//get absolute value of i
		var abs_i int = int(math.Abs(float64(i)))

		//get length, but if 0 we force value to = 1
		length := digitCount(abs_i)
		if length == 0 {
			length++
		}

		//go's equivalent of a while loop
		for length != 1 {
			if abs_i%10 == 5 {
				//break condition
				length = 1
			} else {
				abs_i /= 10
				length--
			}
		}

		if abs_i%10 != 5 {
			count++
		}
	}
	return count
}

func digitCount(number int) int {
	var dig = 0
	for number != 0 {
		number /= 10
		dig++
	}
	return dig
}
