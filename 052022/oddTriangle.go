//Given the triangle of consecutive odd numbers:
//
// 			1
// 	   	     3     5
// 	          7     9    11
// 	      13    15    17    19
//         21    23    25    27    29
//
// Calculate the sum of the numbers in the nth row of this
// triangle (starting at index 1) e.g.: (Input --> Output)
// 1 --> 1
// 2 --> 3 + 5 = 8
// 3 --> 7 + 9 + 11 = 27
//
// Optimization: When I wrote this originally, I misread the prompt and solved it to sum
// all rows before the nth row as well as the nth row. Failing to recognize
// that the sum of each row was n^3, I overcomplicated the solution as I refactored
// to correct my mistake. This is the optimal solution:
package kata

func RowSumOddNumbers(n int) int {
	return n * n * n
}

// This solution is still good, though inefficient.
package kata

func RowSumOddNumbers(n int) int {

	var rowStart int = 1
	var sum int = 0

	for i := 1; i < n; i++ {
		rowStart += 2 * i
	}

	for j := 1; j <= n; j++ {
		sum += rowStart
		rowStart += 2
	}

	return sum
}
