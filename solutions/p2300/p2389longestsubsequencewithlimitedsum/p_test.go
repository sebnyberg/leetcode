package p2389longestsubsequencewithlimitedsum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_answerQueries(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		queries []int
		want    []int
	}{
		{[]int{736411, 184882, 914641, 37925, 214915}, []int{331244, 273144, 118983, 118252, 305688, 718089, 665450}, []int{2, 2, 1, 1, 2, 3, 3}},
		{[]int{4, 5, 2, 1}, []int{3, 10, 21}, []int{2, 3, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, answerQueries(tc.nums, tc.queries))
		})
	}
}

type itemsWithIdx struct {
	items []int
	idx   []int
}

func (l *itemsWithIdx) Less(i, j int) bool { return l.items[i] < l.items[j] }
func (l *itemsWithIdx) Swap(i, j int) {
	l.items[i], l.items[j] = l.items[j], l.items[i]
	l.idx[i], l.idx[j] = l.idx[j], l.idx[i]
}
func (l *itemsWithIdx) Len() int { return len(l.items) }

func answerQueries(nums []int, queries []int) []int {
	n := len(queries)
	itemsWithIdx := itemsWithIdx{
		items: queries,
		idx:   make([]int, n),
	}
	for i := range itemsWithIdx.idx {
		itemsWithIdx.idx[i] = i
	}
	sort.Sort(&itemsWithIdx)
	sort.Ints(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = len(nums)
	}
	var sum int
	var j int
	for i, x := range nums {
		for j < n && sum+x > itemsWithIdx.items[j] {
			res[itemsWithIdx.idx[j]] = i
			j++
		}
		sum += x
		if j == n {
			break
		}
	}

	return res
}
