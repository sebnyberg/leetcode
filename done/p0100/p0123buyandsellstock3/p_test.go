package p0123buyandsellstock3

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"text/tabwriter"

	"github.com/stretchr/testify/require"
)

func dpString(dp [][]int, colname string, rowname string) string {
	var s bytes.Buffer
	w := tabwriter.NewWriter(&s, 0, 0, 1, ' ', 0)

	// Write header
	colIndices := make([]string, len(dp[0]))
	for i := range dp[0] {
		colIndices[i] = strconv.Itoa(i)
	}
	w.Write([]byte(" \t" + colname + "\t" + strings.Join(colIndices, "\t") + "\n"))
	fmt.Println(string(rune(218)))
	w.Write([]byte(rowname + "\t\\\t" + string(strings.Repeat("-\t", len(dp[0]))[:len(dp[0])*2]) + "\n"))

	for i, row := range dp {
		vals := make([]string, len(row))
		for j, entry := range row {
			vals[j] = strconv.Itoa(entry)
		}
		w.Write([]byte(strconv.Itoa(i) + "\t|\t" + strings.Join(vals, "\t") + "\n"))
	}
	w.Flush()
	return s.String()
}

func Test_maxProfit(t *testing.T) {
	for _, tc := range []struct {
		prices []int
		want   int
	}{
		{[]int{1, 4, 2, 7}, 8},
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		{[]int{1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.prices), func(t *testing.T) {
			require.Equal(t, tc.want, maxProfit(tc.prices))
		})
	}
}

// func maxProfit(prices []int) (profit int) {
// 	ntrades := 2
// 	ndays := len(prices)
// 	profits := make([][]int, ntrades)
// 	for i := range profits {
// 		profits[i] = make([]int, ndays)
// 	}

// 	// Fill first row with profits
// 	minVal := prices[0]
// 	for i := 1; i < ndays; i++ {
// 		profits[0][i] = max(profits[0][i-1], prices[i]-minVal)
// 		minVal = min(minVal, prices[i])
// 	}

// 	adjustedPrices := make([]int, ndays)
// 	adjustedPrices[0] = prices[0]

// 	// Solution 1
// 	// for k := 1; k < ntrades; k++ {
// 	// 	for i := 1; i < ndays; i++ {
// 	// 		adjustedPrices[i] = prices[i] - profits[k-1][i-1]
// 	// 		minPriceBeforeToday := adjustedPrices[0]
// 	// 		for j := 1; j < i; j++ {
// 	// 			minPriceBeforeToday = min(minPriceBeforeToday, adjustedPrices[j])
// 	// 		}
// 	// 		profits[k][i] = max(profits[k][i-1], prices[i]-minPriceBeforeToday)
// 	// 	}
// 	// }

// 	// Solution 2 - keep track of minimum values in a list
// 	// minPrices := make([]int, ndays)
// 	// minPrices[0] = prices[0]
// 	// for k := 1; k < ntrades; k++ {
// 	// 	for i := 1; i < ndays; i++ {
// 	// 		minPrices[i] = min(minPrices[i-1], prices[i]-profits[k-1][i-1])
// 	// 		profits[k][i] = max(profits[k][i-1], prices[i]-minPrices[i])
// 	// 	}
// 	// }

// 	// Solution 3 - single minimum value
// 	var minPrice int
// 	for k := 1; k < ntrades; k++ {
// 		minPrice = prices[0]
// 		for i := 1; i < ndays; i++ {
// 			// Deducting the minimum price by the profits the day / trade before,
// 			// we artificially add value to to the current trade
// 			minPrice = min(minPrice, prices[i]-profits[k-1][i-1])
// 			profits[k][i] = max(profits[k][i-1], prices[i]-minPrice)
// 		}
// 	}

// 	return profits[ntrades-1][ndays-1]
// }

func maxProfit(prices []int) (profit int) {
	ntrades := 2
	ndays := len(prices)
	profits := make([]int, ndays)

	var minPrice int
	for k := 0; k < min(ntrades, ndays/2); k++ {
		minPrice = prices[0]
		for i := 1; i < ndays; i++ {
			profits[i], minPrice = max(profits[i-1], prices[i]-minPrice),
				min(minPrice, prices[i]-profits[i])
		}
	}

	return profits[ndays-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
