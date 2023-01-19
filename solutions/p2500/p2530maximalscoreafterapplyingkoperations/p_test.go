package p2530maximalscoreafterapplyingkoperations

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxKelements(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int64
	}{
		{
			[]int{597189039, 57948756, 143524875, 379494516, 862193035, 868775043, 395597119, 275046118, 306907315, 257034002, 476132995, 69495282, 395493151, 354621370, 365510017, 520479568, 219063577, 159958079, 113409455, 170145739, 687892872, 881301934, 723211517, 276655363, 635301113, 440291651, 961908086, 821028930, 821879600, 82879805, 850787822, 547409867, 813461937, 866639644, 512259589, 130847041, 973334294, 114942610, 233744177, 941195642, 888940360, 983125701, 533826303, 726965368, 516401603, 312579605, 182667172, 447853195, 275822190, 338282009},
			62126,
			1,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxKelements(tc.nums, tc.k))
		})
	}
}

func maxKelements(nums []int, k int) int64 {
	count := make(map[int]int)
	for _, x := range nums {
		count[x]++
	}
	var h intHeap
	for x := range count {
		h = append(h, x)
	}
	heap.Init(&h)
	var res int
	for k > 0 {
		x := heap.Pop(&h).(int)
		kk := min(k, count[x])
		delete(count, x)
		k -= kk
		res += kk * x
		next := x / 3
		if x%3 != 0 {
			next++
		}
		if _, exists := count[next]; !exists {
			heap.Push(&h, next)
		}
		count[next] += kk
	}
	return int64(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type intHeap []int

func (h intHeap) Len() int { return len(h) }
func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *intHeap) Push(x interface{}) {
	el := x.(int)
	*h = append(*h, el)
}
func (h *intHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	*h = (*h)[:n-1]
	return el
}
