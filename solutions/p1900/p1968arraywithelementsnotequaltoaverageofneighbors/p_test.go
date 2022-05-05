package p1968arraywithelementsnotequaltoaverageofneighbors

import (
	"fmt"
	"sort"
	"testing"
)

func Test_rearrangeArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 4, 5, 3}},
		{[]int{6, 2, 0, 9, 7}, []int{9, 7, 6, 2, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			res := rearrangeArray(tc.nums)
			if len(res) != len(tc.nums) {
				t.Fatal("result must be same length as input")
			}
			for i := 1; i < len(tc.nums)-1; i++ {
				avg := res[i-1] + res[i+1]
				if avg == res[i]*2 {
					t.Fatalf("failed at %v\nitems: %v\nres: %v\n", i, res[i-1:i+2], res)
				}
			}
		})
	}
}

func rearrangeArray(nums []int) []int {
	// Half the numbers are smaller than the median, half are larger.
	// A combination of two numbers from the 'large' portion cannot have an
	// average of a number from the 'small' portion.
	// Therefore, we can sort nums, then weave numbers from the small and large
	// portion to form a valid result
	sort.Ints(nums)
	half := len(nums) / 2
	res := make([]int, len(nums))
	offset := len(nums) % 2
	for i := 0; i < len(nums)/2; i++ {
		j := i + half + offset
		res[i*2] = nums[i]
		res[i*2+1] = nums[j]
	}
	if len(nums)%2 == 1 {
		res[len(res)-1] = nums[half]
	}
	return res
}
