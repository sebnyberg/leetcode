package p2080rangefrequencyqueries

import (
	"fmt"
	"sort"
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
	valIndices [10001][]int
}

func Constructor(arr []int) RangeFreqQuery {
	f := RangeFreqQuery{}
	for i, el := range arr {
		f.valIndices[el] = append(f.valIndices[el], i)
	}
	return f
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	leftIdx := sort.SearchInts(this.valIndices[value], left)
	rightIdx := sort.SearchInts(this.valIndices[value], right+1)
	return rightIdx - leftIdx
}
