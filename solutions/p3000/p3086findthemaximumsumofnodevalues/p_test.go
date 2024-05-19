package p3086findthemaximumsumofnodevalues

import "math"

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	var sum int
	var changes int
	minDiff := math.MaxInt32
	for _, x := range nums {
		diff := x ^ k - x
		minDiff = min(minDiff, abs(diff))
		if diff > 0 {
			sum += x ^ k
			changes ^= 1
		} else {
			sum += x
		}
	}
	if changes == 1 {
		// odd number of improvements
		// we should pair up the positive improvement with the smallest
		// negative or positive improvement
		sum -= minDiff
	}
	return int64(sum)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
