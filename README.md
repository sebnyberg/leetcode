# Leetcode

Collection of leetcode solutions in Go.

## Permutations

### Lexicographically sorted permutations

[Source](https://stackoverflow.com/questions/31425531/heaps-algorithm-for-permutations)

1. Find the rightmost element which is smaller than the element to its right
2. Swap that element with the smallest element to its right which is larger than it
3. Reverse the part of the permutation to the right of where that element was

### Maximum subarray sum (Kadane's algorithm)

1. For each number
2. If the previous number is greater than zero, add it to the current number
3. If the current number is greater than the maximum sum, it is the new maximum sum

```go
func maxSubArray(nums []int) int {
	n := len(nums)
	m := nums[0]
	for i := 1; i < n; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}
```