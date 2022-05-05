package p0398randompickindex

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	s := Constructor([]int{1, 2, 3, 3, 3})
	res := s.Pick(1)
	require.Equal(t, 0, res)
	var count [5]int
	for i := 0; i < 10000; i++ {
		count[s.Pick(3)]++
	}
	require.InEpsilon(t, 0.333, float64(count[3])/10000, 0.05)
}

type Solution struct {
	numIndices map[int][]int
}

func Constructor(nums []int) Solution {
	s := Solution{
		numIndices: make(map[int][]int, len(nums)/2),
	}
	for i, num := range nums {
		s.numIndices[num] = append(s.numIndices[num], i)
	}
	return s
}

func (this *Solution) Pick(target int) int {
	return this.numIndices[target][rand.Intn(len(this.numIndices[target]))]
}
