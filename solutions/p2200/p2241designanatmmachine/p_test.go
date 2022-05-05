package p2241

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestATM(t *testing.T) {
	a := Constructor()
	a.Deposit([]int{0, 0, 1, 2, 1})
	res := a.Withdraw(600)
	require.Equal(t, []int{0, 0, 1, 0, 1}, res)

	a.Deposit([]int{0, 1, 0, 1, 1})
	res = a.Withdraw(600)
	require.Equal(t, []int{-1}, res)

	res = a.Withdraw(550)
	require.Equal(t, []int{0, 1, 0, 0, 1}, res)
}

var vals [5]int = [5]int{
	20, 50, 100, 200, 500,
}

type ATM struct {
	counts [5]int
}

func Constructor() ATM {
	return ATM{}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, v := range banknotesCount {
		this.counts[i] += v
	}
}

func (this *ATM) Withdraw(amount int) []int {
	d := make([]int, 5)
	for i := 4; i >= 0; i-- {
		d[i] = min(this.counts[i], amount/vals[i])
		amount -= d[i] * vals[i]
	}
	if amount != 0 {
		return []int{-1}
	}
	for i, v := range d {
		this.counts[i] -= v
	}
	return d
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
