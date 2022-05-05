package p2143choosenumbersfromtwoarraysinrange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var a int

func BenchmarkCountSubranges(b *testing.B) {
	nums1 := []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}
	nums2 := []int{42, 87, 55, 37, 63, 21, 67, 88, 98, 100, 8, 51, 33, 44, 43, 66, 52, 9, 41, 36, 47, 4, 70, 13, 19, 16, 47, 39, 88, 29, 97, 7, 14, 23, 59, 43, 7, 13, 92, 91, 63, 15, 85, 9, 3, 66, 76, 57, 85, 95, 76, 53, 5, 44, 28, 5, 41, 58, 27, 54, 18, 64, 73, 76, 43, 42, 29, 25, 99, 22, 100, 48, 27, 60, 39, 11, 87, 3, 37, 25, 87, 86, 70, 41, 25, 28, 16, 97, 80, 53, 62, 29, 15, 2, 21, 87, 55, 4, 64, 100}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		n1 := make([]int, len(nums1))
		n2 := make([]int, len(nums2))
		copy(n1, nums1)
		copy(n2, nums2)
		b.StartTimer()
		a = countSubranges(nums1, nums2)
	}
}

func Test_countSubranges(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
			[]int{42, 87, 55, 37, 63, 21, 67, 88, 98, 100, 8, 51, 33, 44, 43, 66, 52, 9, 41, 36, 47, 4, 70, 13, 19, 16, 47, 39, 88, 29, 97, 7, 14, 23, 59, 43, 7, 13, 92, 91, 63, 15, 85, 9, 3, 66, 76, 57, 85, 95, 76, 53, 5, 44, 28, 5, 41, 58, 27, 54, 18, 64, 73, 76, 43, 42, 29, 25, 99, 22, 100, 48, 27, 60, 39, 11, 87, 3, 37, 25, 87, 86, 70, 41, 25, 28, 16, 97, 80, 53, 62, 29, 15, 2, 21, 87, 55, 4, 64, 100},
			848337464},
		{[]int{77, 22, 29, 91, 52, 37, 41, 93, 99, 32, 56, 3, 74, 38, 26, 68, 6, 16, 98, 52, 24, 83, 89, 0, 2, 92, 37, 50, 9, 46, 45, 60, 63, 60, 65, 47, 78, 94, 50, 58, 46, 77, 8, 68, 37, 80, 7, 88, 2, 37, 83, 87, 25, 8, 85, 2, 100, 36, 49, 11, 18, 79, 38, 46, 73, 42, 46, 50, 85, 41, 37, 60, 6, 2, 73, 83, 26, 29, 86, 100, 74, 41, 91, 12, 4, 19, 0, 30, 38, 14, 100, 31, 92, 83, 76, 41, 74, 29, 57, 82},
			[]int{60, 66, 58, 49, 7, 72, 59, 91, 12, 53, 23, 15, 36, 49, 49, 75, 29, 17, 80, 46, 89, 69, 90, 41, 55, 95, 42, 32, 7, 59, 7, 46, 34, 29, 73, 73, 25, 26, 84, 12, 68, 74, 0, 40, 59, 36, 16, 29, 99, 20, 63, 32, 62, 99, 21, 47, 27, 21, 63, 68, 64, 60, 33, 5, 80, 19, 53, 90, 70, 45, 4, 20, 4, 97, 2, 85, 46, 9, 69, 93, 80, 99, 8, 97, 83, 13, 88, 52, 48, 34, 68, 75, 74, 72, 32, 70, 29, 49, 42, 29},
			894130710,
		},
		{[]int{1, 2, 5}, []int{2, 6, 3}, 3},
		{[]int{0, 1}, []int{1, 0}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, countSubranges(tc.nums1, tc.nums2))
		})
	}
}

const (
	offset = 5100
	size   = 10201
	mod    = 1e9 + 7
)

func countSubranges(nums1 []int, nums2 []int) int {
	var result int
	var prev, cur, empty [size]int

	// maxCanSub and maxCanAdd keeps track of the sums that could be added from
	// remaining elements in the list of numbers
	maxCanSub, maxCanAdd := 0, 0
	for i := range nums1 {
		maxCanAdd += nums1[i]
		maxCanSub += nums2[i]
	}

	// minValid, maxValid limits the range to iterate over in the sums array
	minValid, maxValid := offset, offset

	for i, n1 := range nums1 {
		n2 := -nums2[i]

		minValid = max(offset-maxCanAdd, minValid+n2)
		maxValid = min(offset+maxCanSub, maxValid+n1)
		maxCanAdd -= nums1[i]
		maxCanSub -= nums2[i]

		// Reset current sums array
		cur = empty

		// Add current element (new sum starting in thiselement)
		cur[offset+n1]++
		cur[offset+n2]++

		// Add current number to all sums ending in previous position
		for sum, count := range prev[minValid : maxValid+1] {
			sum += minValid
			cur[sum+n1] = (cur[sum+n1] + count) % mod
			cur[sum+n2] = (cur[sum+n2] + count) % mod
		}

		result = (result + cur[offset]) % mod
		cur, prev = prev, cur
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
