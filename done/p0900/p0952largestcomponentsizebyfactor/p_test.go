package p0952largestcomponentsizebyfactor

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestComponentSize(t *testing.T) {
	for _, tc := range []struct {
		A    []int
		want int
	}{
		{[]int{4, 6, 15, 35}, 4},
		{[]int{20, 50, 9, 63}, 2},
		{[]int{2, 3, 6, 7, 4, 12, 21, 39}, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.A), func(t *testing.T) {
			require.Equal(t, tc.want, largestComponentSize(tc.A))
		})
	}
}

func largestComponentSize(A []int) int {
	dsu := NewDSU(100001)
	for _, n := range A {
		for factor := 2; factor <= int(math.Sqrt(float64(n))); factor++ {
			if n%factor == 0 {
				dsu.union(n, factor)
				dsu.union(n, n/factor)
			}
		}
	}
	groupSize := make(map[int]int)
	for _, n := range A {
		groupSize[dsu.find(n)]++
	}
	maxSize := 1
	for _, size := range groupSize {
		maxSize = max(maxSize, size)
	}

	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return dsu
}

func (d *DSU) find(a int) int {
	if d.parent[a] == a {
		return a
	}
	root := d.find(d.parent[a])
	d.parent[a] = root
	return root
}

func (d *DSU) union(a, b int) {
	a = d.find(a)
	b = d.find(b)
	if a != b {
		if d.size[a] < d.size[b] {
			a, b = b, a
		}
		d.parent[b] = a
		d.size[a] += d.size[b]
	}
}
