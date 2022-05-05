package p2076processrestritedfriendrequests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_friendRequests(t *testing.T) {
	for _, tc := range []struct {
		n                      int
		restrictions, requests [][]int
		want                   []bool
	}{
		{3, [][]int{{0, 1}}, [][]int{{0, 2}, {2, 1}}, []bool{true, false}},
		{3, [][]int{{0, 1}}, [][]int{{1, 2}, {0, 2}}, []bool{true, false}},
		{5, [][]int{{0, 1}, {1, 2}, {2, 3}}, [][]int{{0, 4}, {1, 2}, {3, 1}, {3, 4}}, []bool{true, false, true, false}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, friendRequests(tc.n, tc.restrictions, tc.requests))
		})
	}
}

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	parent := make([]int, n)
	declines := make([]map[int]struct{}, n)
	contains := make([]map[int]struct{}, n)
	for i := range parent {
		parent[i] = i
		declines[i] = make(map[int]struct{})
		contains[i] = map[int]struct{}{i: {}}
	}

	for _, restriction := range restrictions {
		a, b := restriction[0], restriction[1]
		declines[a][b] = struct{}{}
		declines[b][a] = struct{}{}
	}

	// DSU find
	var find func(a int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		return find(parent[a])
	}

	// DSU union
	union := func(a, b int) bool {
		aRoot, bRoot := find(a), find(b)
		for el := range contains[aRoot] {
			if _, exists := declines[bRoot][el]; exists {
				return false
			}
		}
		if aRoot != bRoot {
			parent[aRoot] = bRoot
			for aDecline := range declines[aRoot] {
				declines[bRoot][aDecline] = struct{}{}
			}
			for aContains := range contains[aRoot] {
				contains[bRoot][aContains] = struct{}{}
			}
		}
		return true
	}

	res := make([]bool, len(requests))
	for i := range res {
		a, b := requests[i][0], requests[i][1]
		res[i] = union(a, b)
	}
	return res
}
