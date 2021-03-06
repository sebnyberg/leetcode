# Subarray problem thought process

## Intuition

Consider:

* Whether there is a transformation that would make the problem easier to solve (abs, reduce by k, etc)
* Prefix sums
* If matching is going on, consider DP or KMP
* For next smaller/greater before/after, consider monotonous stacks
* Consider stacks for counts of elements within window
* For hard exercises, sometimes it helps to use binary search
* When dealing with absolute comparisons, examine each individual case, e.g. |a-b|+|c-d| => 12 (4!) cases, of which maybe a subset makes sense in the scenario

### Prefix sums

Prefix sums are useful whenever it is possible to represent the sub-array as a binary operation between two sums.

For example, with a positive sum (+):

```go
n := len(arr)
prefixSum := make([]int, n+1)
for i, num := range arr {
  prefixSum[i+1] = prefixSum[i] + n
}
// Fetch interval between i and j
subArraySum := prefixSum[j] - prefixSum[i-1]
```

The same operation works for any associative operations (mul, div, add, sub, xor).

### Transforming input to work with prefix sums

Consider transforming the input in such a way that it works with prefix sums.

Examples:
* To check whether an array contains any length K sequence greater than a specified average, reduce the value of each element by the required value and check whether the prefix sum is greater than zero.
* (1248) To count odd/even numbers, transform and use prefix sum to count.

### Using stacks

Stacks are very useful for sub-array problems, especially when referring to greatest / smallest element in a sub-array. For example:

* Count number of sub-arrays for which the leftmost number is the minimum value in the subarray
* (907) Count sum of min values of all possible subarrays

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

### Other notes

* Finding the maximum average is the same as finding the maximum sum
* Counting number of possible sub arrays meeting a criterion is equal to multiplying the number of elements to the left and right of the smallest center interval which still meet the criteria.
* Kadane's algorithm can be used to find min sums as well as max sums at the same time
* (918) Maximizing a circular sum that may span from the end to the beginning of an array may be equal to removing the minimum subarray sum from the middle.