//D&C We want to split the array down the middle and calculate the MSS (MaximumSubarraySum) for
//each side. Additionally, we want to consider a solution where the MSS starts in the left side
// and ends in the right side - so we calculate the highest value starting from the left and the highest
// MSS value ending in the right.
// In the example, this would create two slices: [-2,1,-3,4] and [-1,2,1,-5,4]
// the MSS on the left is 4, the MSS on the right is 4, the highest MSS starting from the left is 4 and ending on
// the right is 2 --> so the cross-sectional MSS solution is 6 and is the highest of (4, 4, 6).

package kata

func MaximumSubarraySum(numbers []int) int {

	if len(numbers) == 0 { //Initial handling for empty arrays
		return 0
	} else if len(numbers) == 1 { //base case
		return numbers[0]
	}

	var m int = len(numbers) / 2
	var n int = len(numbers) - m

	//Get leftSlice
	var leftSlice []int = numbers[0:m]
	//Get rightSlice
	var rightSlice []int = numbers[m : n+m]

	//Calculate MSS of left, right, and cross-section of left and right
	var leftMSS int = MaximumSubarraySum(leftSlice)
	var rightMSS int = MaximumSubarraySum(rightSlice)
	var crossMSS int = maxCrossing(numbers, 0, len(numbers)-1, m)

	//Exception handling for when the entire list is negative:
	if max(leftMSS, rightMSS, crossMSS) < 0 {
		return 0
	}

	//Return greatest of the three MSS conditions
	return max(leftMSS, rightMSS, crossMSS)

}

//Helper function to find the MSS which starts in the left half and ends in the right half of the original array
func maxCrossing(nums []int, left int, right int, mid int) int {

	var leftSum int = -999999
	var rightSum int = -999999
	var sum int = 0

	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	sum = 0

	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}

	if leftSum < 0 {
		leftSum = 0
	} else if rightSum < 0 {
		rightSum = 0
	}
	return leftSum + rightSum
}

//Helper function to get the maximum of any number of ints
func max(values ...int) int {
	var maxVal int = values[0]
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
