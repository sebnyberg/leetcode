package p1675minimizedevinarr

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func mustRead(fpath string) []int {
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatalln(err)
	}
	res := make([]int, 0)
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	for _, s := range strings.Split(string(content), ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, i)
	}
	return res
}

func Test_minimumDeviation(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{399, 908, 648, 357, 693, 502, 331, 649, 596, 698}, 315},
		{[]int{4, 1, 5, 20, 3}, 3},
		{[]int{2, 10, 8}, 3},
		{[]int{1, 2, 3, 4}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumDeviation(tc.nums))
		})
	}
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] > h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	ans := old[n-1]
	*h = old[:n-1]
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minimumDeviation(nums []int) int {
	n := len(nums)
	pq := make(Heap, 0, n)

	// Put the maximum possible numbers into the heap
	minN := int(1e9)
	for _, num := range nums {
		if num%2 == 1 {
			num *= 2
		}
		minN = min(minN, num)
		heap.Push(&pq, num)
	}

	ans := int(1e9)
	for {
		maxN := heap.Pop(&pq).(int)
		ans = min(ans, maxN-minN)
		if maxN%2 == 1 {
			break
		}
		heap.Push(&pq, maxN/2)
		minN = min(minN, maxN/2)
	}
	return ans
}
