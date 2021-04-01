package p0399evaluatedivision

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calcEquation(t *testing.T) {
	for _, tc := range []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		// {[][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}, []float64{6, 0.5, -1, -1, -1}},
		// {[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}, []float64{0.5, 2, -1, -1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.equations), func(t *testing.T) {
			require.Equal(t, tc.want, calcEquation(tc.equations, tc.values, tc.queries))
		})
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	dsu := NewDSU(len(equations) * 2)
	eqIndices := make(map[string]int)
	var eqIdx int
	for i, eq := range equations {
		if _, exists := eqIndices[eq[0]]; !exists {
			eqIndices[eq[0]] = eqIdx
			eqIdx++
		}
		if _, exists := eqIndices[eq[1]]; !exists {
			eqIndices[eq[1]] = eqIdx
			eqIdx++
		}
		r1 := dsu.FindRoot(eqIndices[eq[0]])
		r2 := dsu.FindRoot(eqIndices[eq[1]])
		dsu.parent[r1] = r2
		dsu.dist[r1] = dsu.dist[eqIndices[eq[0]]] * values[i] / dsu.dist[eqIndices[eq[1]]]
	}

	// resp := make([]float64, 0)
	// for _, query := range queries {
	// 	i, j := dsu.Find(eqIndices[query[0]]), dsu.Find(eqIndices[query[1]])
	// 	if i != j {
	// 		resp = append(resp, -1)
	// 	}
	// 	// there is a shared root
	// }
	return nil
}

type DSU struct {
	parent []int
	dist   []float64 // distance from i to parent[i]
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		dist:   make([]float64, n),
	}

	for i := range dsu.parent {
		dsu.parent[i] = i
		dsu.dist[i] = 1 // the distance of any node to itself is 1
	}

	return dsu
}

func (d *DSU) FindRoot(x int) int {
	if d.parent[x] == x {
		return x
	}
	lastp := d.parent[x]
	p := d.FindRoot(lastp)
	d.parent[x] = p
	d.dist[x] = d.dist[x] * d.dist[lastp]
	return p
}
