package p2483minimumpenaltyforashop

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_bestClosingTime(t *testing.T) {
	for i, tc := range []struct {
		customers string
		want      int
	}{
		{"YYNY", 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, bestClosingTime(tc.customers))
		})
	}
}

func bestClosingTime(customers string) int {
	n := len(customers)
	left := make([]int, n+1)
	for i := range customers {
		left[i+1] = left[i]
		if customers[i] == 'N' {
			left[i+1]++
		}
	}
	right := make([]int, n+1)
	for i := len(customers) - 1; i >= 0; i-- {
		right[i] = right[i+1]
		if customers[i] == 'Y' {
			right[i]++
		}
	}
	minCost := math.MaxInt32
	var res int
	for i := 0; i <= len(customers); i++ {
		if d := left[i] + right[i]; d < minCost {
			minCost = d
			res = i
		}
		minCost = min(minCost, left[i]+right[i])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
