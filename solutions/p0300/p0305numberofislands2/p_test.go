package p0305numberofislands2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numIslands2(t *testing.T) {
	for _, tc := range []struct {
		m         int
		n         int
		positions [][]int
		want      []int
	}{
		{3, 3, [][]int{{0, 1}, {1, 2}, {2, 1}, {1, 0}, {0, 2}, {0, 0}, {1, 1}}, []int{1, 2, 3, 4, 3, 2, 1}},
		{3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}, []int{1, 1, 2, 3}},
		{1, 1, [][]int{{0, 0}}, []int{1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.m), func(t *testing.T) {
			require.Equal(t, tc.want, numIslands2(tc.m, tc.n, tc.positions))
		})
	}
}

func numIslands2(m int, n int, positions [][]int) []int {
	res := make([]int, len(positions))
	dsu := NewDSU(m, n)
	for i, pos := range positions {
		dsu.addLand(pos[0], pos[1])
		res[i] = dsu.islands
	}

	return res
}

type DSU struct {
	parent  []int
	land    []bool
	m       int
	n       int
	islands int
}

func NewDSU(m, n int) DSU {
	dsu := DSU{
		parent: make([]int, m*n),
		land:   make([]bool, m*n),
		m:      m,
		n:      n,
	}
	for i := range dsu.parent {
		dsu.parent[i] = i
	}
	return dsu
}

func (d *DSU) find(a int) int {
	if d.parent[a] == a {
		return a
	}
	// path compression
	root := d.find(d.parent[a])
	d.parent[a] = root
	return root
}

func (d *DSU) addLand(i, j int) {
	a := i*d.n + j
	if d.land[a] {
		return
	}
	d.land[a] = true
	d.islands++

	for _, near := range [][2]int{{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1}} {
		i, j := near[0], near[1]
		if i < 0 || j < 0 || i >= d.m || j >= d.n {
			continue
		}
		b := i*d.n + j
		if !d.land[b] {
			continue
		}
		ra := d.find(a)
		rb := d.find(b)
		if ra != rb {
			d.islands--
		}
		d.parent[ra] = rb
	}
}
