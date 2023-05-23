package p1434numberofwaystoweardifferenthatstoeachother

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_numberWays(t *testing.T) {
	for i, tc := range []struct {
		hats [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[3,4],[4,5],[5]]"),
			1,
		},
		{
			leetcode.ParseMatrix("[[3,5,1],[3,5]]"),
			4,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numberWays(tc.hats))
		})
	}
}

const mod = 1e9 + 7

func numberWays(hats [][]int) int {
	// When a hat is only preffered by a single person, the choice is simple.
	// When there are no more such choices, then a hat must be picked, seemingly
	// at random. We could memoize the number of possible permutations of
	// assignments given that the first N people have chosen their hats, and a
	// bitmask of remaining, available hats.
	//
	// Ok, I tried it. It was too slow. Let's invert the problem from trying
	// hats per person to trying person per hat. Change from hats[i] = hats for
	// person, to person[i] = persons for hat
	var persons [][]int
	var hatCount int
	hatIdx := map[int]int{}
	for i := range hats {
		for _, x := range hats[i] {
			if _, exists := hatIdx[x]; !exists {
				hatIdx[x] = hatCount
				hatCount++
			}
			x = hatIdx[x]
			if x >= len(persons) {
				persons = append(persons, make([][]int, (x+1)-len(persons))...)
			}
			persons[x] = append(persons[x], i)
		}
	}
	mem := make(map[[2]int]int)
	res := dfs(mem, persons, 0, 0, len(hats))
	return res
}

func dfs(mem map[[2]int]int, persons [][]int, i, bm, n int) int {
	if bm == (1<<n)-1 {
		// Everyone has a hat! Good stuff.
		return 1
	}
	if i == len(persons) {
		return 0
	}
	key := [2]int{i, bm}
	if v, exists := mem[key]; exists {
		return v
	}
	// Try to skip this hat, and to assign it to each person
	res := dfs(mem, persons, i+1, bm, n)
	for _, p := range persons[i] {
		if bm&(1<<p) > 0 {
			continue
		}
		res = (res + dfs(mem, persons, i+1, bm|(1<<p), n)) % mod
	}
	mem[key] = res
	return mem[key]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
