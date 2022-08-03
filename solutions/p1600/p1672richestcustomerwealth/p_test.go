package p1672richestcustomerwealth

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumWealth(t *testing.T) {
	for _, tc := range []struct {
		accounts [][]int
		want     int
	}{
		{leetcode.ParseMatrix("[[1,2,3],[3,2,1]]"), 6},
		{leetcode.ParseMatrix("[[1,5],[7,3],[3,5]]"), 10},
		{leetcode.ParseMatrix("[[2,8,7],[7,1,3],[1,9,5]]"), 17},
	} {
		t.Run(fmt.Sprintf("%+v", tc.accounts), func(t *testing.T) {
			require.Equal(t, tc.want, maximumWealth(tc.accounts))
		})
	}
}

func maximumWealth(accounts [][]int) int {
	var maxWealth int
	for _, account := range accounts {
		var sum int
		for _, money := range account {
			sum += money
		}
		if sum > maxWealth {
			maxWealth = sum
		}
	}
	return maxWealth
}
