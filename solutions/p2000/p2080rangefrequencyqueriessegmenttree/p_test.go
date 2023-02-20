package p2080rangefrequencyqueries

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RangeFreqQuery(t *testing.T) {
	for _, tc := range []struct {
		init   []int
		inputs [][]int
		want   []int
	}{
		{[]int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56}, [][]int{{1, 2, 4}, {0, 11, 33}}, []int{1, 2}},
		{[]int{1, 1, 1, 2, 2}, [][]int{{0, 1, 2}, {0, 2, 1}, {3, 3, 2}, {2, 2, 1}}, []int{0, 3, 1, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.init), func(t *testing.T) {
			rfq := Constructor(tc.init)
			for i, input := range tc.inputs {
				require.Equal(t, tc.want[i], rfq.Query(input[0], input[1], input[2]))
			}
		})
	}
}

type RangeFreqQuery struct {
	tree []map[int]int
	n    int
}

func Constructor(arr []int) RangeFreqQuery {
	n := len(arr)
	for bits.OnesCount(uint(n)) != 1 {
		n++
	}
	var r RangeFreqQuery
	r.tree = make([]map[int]int, n*2)
	for i := range arr {
		r.tree[n+i] = map[int]int{arr[i]: 1}
	}
	for i := n - 1; i >= 1; i-- {
		a := r.tree[i*2]
		b := r.tree[i*2+1]
		r.tree[i] = make(map[int]int, len(a)+len(b))
		for k, v := range a {
			r.tree[i][k] += v
		}
		for k, v := range b {
			r.tree[i][k] += v
		}
	}
	r.n = n
	return r
}

func (this *RangeFreqQuery) q(i, lo, hi, qlo, qhi, val int) int {
	if lo >= qlo && hi <= qhi {
		return this.tree[i][val]
	}
	if lo >= qhi || hi <= qlo {
		return 0
	}
	return this.q(2*i, lo, lo+(hi-lo)/2, qlo, qhi, val) +
		this.q(2*i+1, lo+(hi-lo)/2, hi, qlo, qhi, val)
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	return this.q(1, 0, this.n, left, right+1, value)
}
