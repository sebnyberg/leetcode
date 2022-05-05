package p2163minimumdifferenceinsumsafterremovalofelements

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumDifference(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{7, 2, 5, 9, 2, 1, 6, 7, 4}, -12},
		{[]int{3, 1, 2}, -1},
		{[]int{7, 9, 5, 8, 1, 3}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumDifference(tc.nums))
		})
	}
}

func minimumDifference(nums []int) int64 {
	m := len(nums)
	n := m / 3

	// Goal is to minimize left side, and maximize right side
	// So we have a max heap on left side, and min on right side

	// Idea:
	// Go to 2*n, add min items to a heap. The
	// Load [2*n, 3*n) into max heap
	// Move cursor from 2*n-1 to n,
	var rightSum int
	rightMin := MinHeap{}
	for i := len(nums) - 1; i >= 2*n; i-- {
		rightSum += nums[i]
		rightMin = append(rightMin, nums[i])
	}
	heap.Init(&rightMin)
	var leftSum int
	leftMin := MaxHeap{}
	for i := 0; i < 2*n; i++ {
		heap.Push(&leftMin, nums[i])
		if len(leftMin) > n {
			// Pop smallest, add to sum
			leftSum += heap.Pop(&leftMin).(int)
		}
	}
	minDiff := leftSum - rightSum
	// Now move the cursor from 2*n to n, checking the min difference
	for i := 2*n - 1; i >= n; i-- {
		// Add to possible minimum values to remove from right sum
		heap.Push(&rightMin, nums[i])
		rightSum += nums[i]
		// We can remove the mimimum value from right side
		minRightVal := heap.Pop(&rightMin).(int)
		rightSum -= minRightVal
		if nums[i] < leftMin[0] {
			leftSum -= nums[i]
			leftSum += heap.Pop(&leftMin).(int)
		} else {
			// Remove from heap
			for j := range leftMin {
				if leftMin[j] == nums[i] {
					heap.Remove(&leftMin, j)
					break
				}
			}
		}
		minDiff = min(minDiff, leftSum-rightSum)
	}
	return int64(minDiff)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
