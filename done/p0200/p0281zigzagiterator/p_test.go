package p0281zigzagiterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ZigzagIterator(t *testing.T) {
	for _, tc := range []struct {
		v    [2][]int
		want []int
	}{
		{[2][]int{{1, 2}, {3, 4, 5, 6}}, []int{1, 3, 2, 4, 5, 6}},
		{[2][]int{{1}, {}}, []int{1}},
		{[2][]int{{}, {1}}, []int{1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.v), func(t *testing.T) {
			z := Constructor(tc.v[0], tc.v[1])
			res := make([]int, 0)
			for z.hasNext() {
				res = append(res, z.next())
			}
			require.Equal(t, tc.want, res)
		})
	}
}

type ZigzagIterator struct {
	arrs      [2][]int
	sharedLen int
	nTotal    int
	biggest   int
	pos       int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	n1 := len(v1)
	n2 := len(v2)
	var biggest int
	if n2 > n1 {
		biggest = 1
	}
	var sharedLen int
	if n1 <= n2 {
		sharedLen = n1 * 2
	} else {
		sharedLen = n2 * 2
	}
	return &ZigzagIterator{
		arrs:      [2][]int{v1, v2},
		nTotal:    n1 + n2,
		sharedLen: sharedLen,
		biggest:   biggest,
		pos:       0,
	}
}

func (this *ZigzagIterator) next() int {
	var res int
	if this.pos < this.sharedLen {
		res = this.arrs[this.pos%2][this.pos/2]
	} else {
		res = this.arrs[this.biggest][this.pos-(this.sharedLen/2)]
	}
	this.pos++
	return res
}

func (this *ZigzagIterator) hasNext() bool {
	return this.pos < this.nTotal
}
