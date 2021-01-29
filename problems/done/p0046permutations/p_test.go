package p0046permutations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_permute(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want [][]int
	}{
		{[]int{0, 1}, [][]int{{0, 1}, {1, 0}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.EqualValues(t, tc.want, permute(tc.nums))
		})
	}
}

func permute(nums []int) [][]int {
	n := len(nums)
	nperms := 1
	for i := 2; i <= n; i++ {
		nperms *= i
	}

	res := make([][]int, nperms)
	p := NewPermutator(n)
	i := 0
	for p.Permutate(func(i, j int) { nums[i], nums[j] = nums[j], nums[i] }) {
		res[i] = make([]int, n)
		copy(res[i], nums)
		i++
	}
	return res
}

type Permutator struct {
	permIndex int
	nperms    int
	// t is the number of swaps per number index
	t []int
	// c is the number of performed swaps per number index
	c []int
}

// Create a new Permutor which permutates an array of length n
func NewPermutator(n int) *Permutator {
	p := Permutator{
		t:      make([]int, n),
		c:      make([]int, n),
		nperms: 1,
	}
	for i := 0; i < n; i++ {
		p.t[i] = i + 1
		p.nperms *= i + 1
	}
	return &p
}

// Permutate permutates the provided array and returns false
// if no more permutations can be done.
func (p *Permutator) Permutate(swapFn func(i, j int)) bool {
	if p.permIndex == 0 {
		p.permIndex++
		return true
	}
	if p.permIndex >= p.nperms {
		return false
	}

	i := 0
	for p.c[i] >= p.t[i] {
		p.c[i] = 0
		i++
	}

	start, end := 0, i+1
	if i%2 == 0 {
		start = p.c[i]
	}
	p.c[i]++

	swapFn(start, end)
	p.permIndex++

	return true
}
