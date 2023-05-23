package p1449formlargestintegerwithdigitsthatadduptotarget

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestNumber(t *testing.T) {
	for i, tc := range []struct {
		cost   []int
		target int
		want   string
	}{
		{[]int{7, 6, 5, 5, 5, 6, 8, 7, 8}, 12, "85"},
		{[]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 9, "7772"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, largestNumber(tc.cost, tc.target))
		})
	}
}

func largestNumber(cost []int, target int) string {
	// Given a set of costs, find the maximum number of combinations of costs
	// (can be repeated) such that the sum is equal to target. It doesn't
	// actually matter what number it is - more is better and if two digits have
	// the same cost, we can simply choose the highest-value alternative.
	//
	// I was very tired when I wrote this - this is a very inefficient
	// implementation. Please don't use it.
	uniq := []int{}
	for _, c := range cost {
		uniq = append(uniq, c)
	}
	sort.Ints(uniq)
	var j int
	for i := range uniq {
		if uniq[i] != uniq[j] {
			j++
		}
		uniq[j] = uniq[i]
	}
	uniq = uniq[:j+1]
	digitForCost := make([]byte, len(uniq))
	for c := '9'; c >= '1'; c-- {
		for i, v := range uniq {
			x := cost[c-'1']
			if v == x && digitForCost[i] == 0 {
				digitForCost[i] = byte(c)
			}
		}
	}

	mem := make(map[[2]int]*string)
	res := dp(mem, uniq, digitForCost, 0, target)
	if res == nil {
		return "0"
	}
	return *res
}

var empty = ""

func dp(mem map[[2]int]*string, uniq []int, digitForCost []byte, i, target int) *string {
	if target < 0 {
		return nil
	}
	if target == 0 {
		return &empty
	}
	key := [2]int{i, target}
	if v, exists := mem[key]; exists {
		return v
	}

	// Try each number
	var res *string
	for j := i; j < len(uniq); j++ {
		a := dp(mem, uniq, digitForCost, j, target-uniq[j])
		if a == nil {
			continue
		}
		bs := append([]byte(*a), digitForCost[j])
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] > bs[j]
		})
		s := string(bs)
		if res == nil || len(s) > len(*res) || (len(s) == len(*res) && s > *res) {
			res = &s
		}
	}
	mem[key] = res
	return res
}
