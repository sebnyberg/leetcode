package p0528randompickwithweight

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	s := Constructor([]int{1, 3})
	var res [2]int
	for i := 0; i < 100; i++ {
		res[s.PickIndex()]++
	}
	require.Equal(t, 0, res[0])
}

type Solution struct {
	presum []int
	tot    int
}

func Constructor(w []int) Solution {
	s := Solution{
		presum: make([]int, len(w)+1),
	}
	s.presum[0] = 0
	for i, x := range w {
		s.presum[i+1] = s.presum[i] + x
	}
	s.tot = s.presum[len(s.presum)-1]
	return s
}

func (this *Solution) PickIndex() int {
	x := 1 + rand.Intn(this.tot)
	idx := sort.SearchInts(this.presum, x) - 1
	return idx
}
