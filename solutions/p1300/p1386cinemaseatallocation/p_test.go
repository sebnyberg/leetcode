package p1386cinemaseatallocation

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_maxNumberOfFamilies(t *testing.T) {
	for i, tc := range []struct {
		n             int
		reservedSeats [][]int
		want          int
	}{
		{
			3,
			leetcode.ParseMatrix("[[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]"),
			4,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxNumberOfFamilies(tc.n, tc.reservedSeats))
		})
	}
}

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	sort.Slice(reservedSeats, func(i, j int) bool {
		a := reservedSeats[i]
		b := reservedSeats[j]
		return a[0] < b[0]
	})
	prev := 0
	var res int
	var j int
	for j < len(reservedSeats) {
		res += 2 * (reservedSeats[j][0] - prev - 1)
		prev += (reservedSeats[j][0] - (prev + 1)) * 2
		var reserved [10]int
		currRow := reservedSeats[j][0]
		prev = reservedSeats[j][0]
		for j < len(reservedSeats) && reservedSeats[j][0] == currRow {
			reserved[reservedSeats[j][1]-1] = 1
			j++
		}
		var presum [11]int
		for i := 0; i < len(reserved); i++ {
			presum[i+1] = presum[i] + reserved[i]
		}
		a := presum[5]-presum[1] == 0
		b := presum[9]-presum[5] == 0
		c := presum[7]-presum[3] == 0
		if a || b {
			if a && b {
				res += 2
			} else {
				res++
			}
			continue
		} else if c {
			res++
		}
	}
	res += 2 * (n - prev)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
