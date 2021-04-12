package p1792maxaveragepassratio

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxAverageRatio(t *testing.T) {
	for _, tc := range []struct {
		classes       [][]int
		extraStudents int
		want          float64
	}{
		{[][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, 4, 0.53485},
		{[][]int{{1, 2}, {3, 5}, {2, 2}}, 2, 0.78333},
	} {
		t.Run(fmt.Sprintf("%+v", tc.classes), func(t *testing.T) {
			require.InEpsilon(t, tc.want, maxAverageRatio(tc.classes, tc.extraStudents), 0.001)
		})
	}
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := make(MinHeap, 0)
	for _, class := range classes {
		pass, total := class[0], class[1]
		h = append(h, &Class{pass, total, profit(pass, total)})
	}
	heap.Init(&h)

	for extraStudents > 0 && h.Len() > 0 {
		h[0].pass++
		h[0].total++
		h[0].profit = profit(h[0].pass, h[0].total)
		heap.Fix(&h, 0)
		extraStudents--
	}

	var value float64
	var count float64
	for _, class := range h {
		value += float64(class.pass) / float64(class.total)
		count++
	}
	res := value / count

	return res
}

// profit of increasing number of students by one
func profit(pass, total int) float64 {
	return (float64(pass+1) / float64(total+1)) - (float64(pass) / float64(total))
}

type Class struct {
	pass, total int
	profit      float64
}

type MinHeap []*Class

func (h *MinHeap) Len() int { return len(*h) }
func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].profit > (*h)[j].profit
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Class))
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}
