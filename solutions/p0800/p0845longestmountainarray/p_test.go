package p0845longestmountainarray

func longestMountain(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}

	// left[i] = length of increasing subarray ending in i
	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	// Any combination of an increasing and decreasing subarray where
	// both subarrays are of length >= 2 must be a valid mountain range
	var maxLength int
	rightLen := 1
	for i := n - 2; i >= 1; i-- {
		if arr[i] <= arr[i+1] {
			rightLen = 1
			continue
		}
		rightLen++
		if left[i] > 1 {
			maxLength = max(maxLength, left[i]+rightLen-1)
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
