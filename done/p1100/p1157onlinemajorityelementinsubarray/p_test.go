package p1157onlinemajorityelementinsubarray

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMajorityChecker(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		m := Constructor([]int{1, 1, 2, 2, 1, 1})
		res := m.Query(0, 5, 4)
		require.Equal(t, 1, res)
		res = m.Query(0, 3, 3)
		require.Equal(t, -1, res)
		res = m.Query(2, 3, 2)
		require.Equal(t, 2, res)
	})
}

type MajorityChecker struct {
	arr map[int][]int
}

func Constructor(arr []int) MajorityChecker {
	m := MajorityChecker{arr: make(map[int][]int)}
	for i, n := range arr {
		m.arr[n] = append(m.arr[n], i)
	}
	return m
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	for num, indices := range this.arr {
		if len(indices) < threshold {
			continue
		}
		// Check if provided interval contains threshold numbers
		lower := sort.SearchInts(indices, left)
		upper := sort.SearchInts(indices, right+1)
		if upper-lower >= threshold {
			return num
		}
	}
	return -1
}
