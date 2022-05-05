package p0465optimalaccountbalancing

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minTransfers(t *testing.T) {
	for _, tc := range []struct {
		transactions [][]int
		want         int
	}{
		{[][]int{{0, 1, 10}, {2, 0, 5}}, 2},
		{[][]int{{0, 1, 10}, {1, 0, 1}, {1, 2, 5}, {2, 0, 5}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.transactions), func(t *testing.T) {
			require.Equal(t, tc.want, minTransfers(tc.transactions))
		})
	}
}

func minTransfers(transactions [][]int) int {
	// There are only up to 8 transactions, so potentially an exponential solution
	// with some smart choices is good enough.
	//
	// At first I thought of this as a graph, but it's not really. Nothing stops
	// a person from giving money to someone with whom they do not have a
	// personal debt.
	//
	var idx [21]int // index compression so that bitset can be used
	for i := range idx {
		idx[i] = -1
	}
	balance := make([]int, 0)
	for _, t := range transactions {
		from, to, amt := t[0], t[1], t[2]
		if idx[from] == -1 {
			idx[from] = len(balance)
			balance = append(balance, 0)
		}
		if idx[to] == -1 {
			idx[to] = len(balance)
			balance = append(balance, 0)
		}
		balance[idx[from]] -= amt
		balance[idx[to]] += amt
	}

	// Store only those who have debt
	debt := make([]int, 0)
	for _, bal := range balance {
		if bal != 0 {
			debt = append(debt, bal)
		}
	}
	res := dfs(debt, 0)
	return res
}

func dfs(balance []int, i int) int {
	for i < len(balance) && balance[i] == 0 {
		i++
	}
	if i == len(balance) {
		return 0
	}

	res := math.MaxInt32
	for j := i + 1; j < len(balance); j++ {
		if balance[i]*balance[j] > 0 {
			continue
		}
		balance[j] += balance[i]
		res = min(res, 1+dfs(balance, i+1))
		balance[j] -= balance[i]
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
