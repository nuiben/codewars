package kata

import (
	"fmt"
	"math"
)

func DontGiveMeFive(start int, end int) int {

	count := 0
	//if we have a number greater than 49,
	//we start running into numbers which have the beginning number as a '5'
	//ex: 50, 500, 5000, 50000
	//to account for this, we need to perform a log10 function on the number
	//in the case of 50's, we do this once and perform the Abs(%10) to see if it
	//contains a 5.

	for i := start; i <= end; i++ {
		fmt.Print("i: ", i, ", ")
		//get absolute value of i
		var abs_i int = int(math.Abs(float64(i)))
		//fmt.Print("abs_i: ",abs_i)
		//is |i| > 50?
		//if so, this simple for-loop will work
		//get length
		length := digitCount(abs_i)
		if length == 0 {
			length++
		}
		fmt.Print("length: ", length, ", ")
		for length != 1 {
			if abs_i%10 == 5 {
				fmt.Println("Check")
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
