package p0621taskscheduler

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_leastInterval(t *testing.T) {
	for _, tc := range []struct {
		tasks []byte
		n     int
		want  int
	}{
		{[]byte("AAAAAABCDEFG"), 2, 16},
		{[]byte("AAABBB"), 2, 8},
		{[]byte("AAABBB"), 0, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tasks), func(t *testing.T) {
			require.Equal(t, tc.want, leastInterval(tc.tasks, tc.n))
		})
	}
}

func leastInterval(tasks []byte, n int) int {
	var taskCount [26]int
	for _, task := range tasks {
		taskCount[task-'A']++
	}
	h := make(MaxHeap, 0)
	for b, count := range taskCount {
		if count > 0 {
			h = append(h, item{
				name:  byte(b),
				count: count,
				time:  -1,
			})
		}
	}
	heap.Init(&h)
	q := []item{}
	t := 1
	for len(q) > 0 || len(h) > 0 {
		// Clear queue of items that are available
		for len(q) > 0 && q[0].time < t {
			heap.Push(&h, q[0])
			q = q[1:]
		}

		// Pick most voluminous item
		if len(h) > 0 {
			it := heap.Pop(&h).(item)
			it.count--
			if it.count > 0 {
				it.time = t + n
				q = append(q, it)
			}
		}

		t++
	}
	return t - 1
}

type item struct {
	name  byte
	count int
	time  int
}

type MaxHeap []item

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(item))
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
