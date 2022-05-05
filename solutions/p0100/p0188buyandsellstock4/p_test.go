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
		k      int
		prices []int
		want   int
	}{
		{2, []int{2, 4, 1}, 2},
		{2, []int{3, 2, 6, 5, 0, 3}, 7},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.k, tc.prices), func(t *testing.T) {
			require.Equal(t, tc.want, maxProfit(tc.k, tc.prices))
		})
	}
}

func maxProfit(k int, prices []int) (profit int) {
	if len(prices) <= 1 {
		return 0
	}
	ndays := len(prices)
	profits := make([]int, ndays)
	ntrades := k

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
