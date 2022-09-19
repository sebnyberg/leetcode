package p2412minimumnumberrequiredbeforetransactions

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minimumMoney(t *testing.T) {
	for _, tc := range []struct {
		transactions [][]int
		want         int64
	}{
		{
			leetcode.ParseMatrix("[[6,5],[8,5],[9,0],[10,1],[4,10],[3,6],[0,5]]"),
			27,
		},
		{
			leetcode.ParseMatrix("[[7,2],[0,10],[5,0],[4,1],[5,8],[5,9]]"),
			18,
		},
		{
			leetcode.ParseMatrix("[[2,1],[5,0],[4,2]]"),
			10,
		},
		{
			leetcode.ParseMatrix("[[3,0],[0,3]]"),
			3,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.transactions), func(t *testing.T) {
			require.Equal(t, tc.want, minimumMoney(tc.transactions))
		})
	}
}

func minimumMoney(transactions [][]int) int64 {
	var minCost int
	var maxCashback int
	var maxCostWithPositiveDelta int
	var money int
	for _, t := range transactions {
		cost, cashback := t[0], t[1]
		minCost = max(minCost, cost)
		if cashback >= cost {
			maxCostWithPositiveDelta = max(maxCostWithPositiveDelta, cost)
			continue
		}
		maxCashback = max(maxCashback, cashback)
		money += cost - cashback
	}
	money = max(money+maxCashback, money+maxCostWithPositiveDelta)
	res := int64(max(minCost, money))
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
