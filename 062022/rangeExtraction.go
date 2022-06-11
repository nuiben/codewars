package kata

// Import libraries
import (
	"strconv"
)

func Solution(list []int) string {
	// After passing through the helper function, we need to place the leftover elements into the result
	var empty_stack = []int{}
	var result, stack = rangeMaker(list, empty_stack)
	// If there is one element in the stack, we return the result:
	if len(stack) < 2 {
		return result
		// If there are 2 elements leftover, we append the first one as a independent integer
	} else if len(stack) == 2 {
		return strconv.Itoa(stack[0]) + "," + result
		// If there are more than 2 elements, we append the first one as an interval
	} else {
		return strconv.Itoa(stack[0]) + "-" + result
	}

}

// RangeMaker Function:
// Parse in 2 integer arrays, one as the list and one as the empty stack
// Return a result string and the stack after extracting ranges
func rangeMaker(list []int, stack []int) (string, []int) {
	// Base Case, if there is only 1 element in the list, we simply return the first element of the list as stack
	if len(list) == 1 {
		return strconv.Itoa(list[0]), []int{list[0]}
	}

	// Recursion: process the list from the 1st element each time so that the list is shrinking
	var result, newStack = rangeMaker(list[1:], stack)

	// If the current element is strictly equal to the top of stack minus 1, i.e within that range
	if list[0] == newStack[0]-1 {
		// We add the element to the stack
		newStack = append([]int{list[0]}, newStack...)
		return result, newStack
	} else {
		// If the current element is not the in the range
		// Then we basically add the stack to the result and empty the stack
		if len(newStack) < 2 {
			result = strconv.Itoa(list[0]) + "," + result
		} else if len(newStack) == 2 {
			result = strconv.Itoa(list[0]) + "," + strconv.Itoa(newStack[0]) + "," + result
		} else {
			result = strconv.Itoa(list[0]) + "," + strconv.Itoa(newStack[0]) + "-" + result
		}
	}
	// Then new stack would be the first element of the list
	newStack = []int{list[0]}
	return result, newStack
}
