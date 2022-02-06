package p2165smallestvalueoftherearrangednumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestNumber(t *testing.T) {
	for _, tc := range []struct {
		num  int64
		want int64
	}{
		{310, 103},
		{-7605, -7650},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, smallestNumber(tc.num))
		})
	}
}

func smallestNumber(num int64) int64 {
	var numCount [10]int
	var neg bool
	if num < 0 {
		neg = true
		num = -num
	}
	for x := num; x > 0; x /= 10 {
		numCount[x%10]++
	}
	if num == 0 {
		return 0
	}
	// If num is negative, we want to maximize the number
	if neg {
		var res int
		for num := 9; num >= 0; num-- {
			for i := 0; i < numCount[num]; i++ {
				res *= 10
				res += num
			}
		}
		return int64(-res)
	}
	// Num is positive, add first the smallest non-zero number
	var res int
	for num := 1; num <= 9; num++ {
		if numCount[num] > 0 {
			res += num
			numCount[num]--
			break
		}
	}
	// Then add numbers from small to large
	for num := 0; num <= 9; num++ {
		for i := 0; i < numCount[num]; i++ {
			res *= 10
			res += num
		}
	}
	return int64(res)
}
