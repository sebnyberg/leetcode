package p0166fractiontorecurringdecimal

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fractionToDecimal(t *testing.T) {
	for _, tc := range []struct {
		numerator   int
		denominator int
		want        string
	}{
		// {1, 2, "0.5"},
		// {1, 333, "0.(003)"},
		{1, 6, "0.1(6)"},
		{4, 333, "0.(012)"},
		{1, 5, "0.2"},
		{-1, 5, "-0.2"},
		{1, -5, "-0.2"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.numerator), func(t *testing.T) {
			require.Equal(t, tc.want, fractionToDecimal(tc.numerator, tc.denominator))
		})
	}
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	res := make([]byte, 0)
	switch {
	case numerator < 0 && denominator < 0:
		numerator = -numerator
		denominator = -denominator
	case numerator < 0 && denominator >= 0:
		numerator = -numerator
		res = append(res, '-')
	case numerator >= 0 && denominator < 0:
		denominator = -denominator
		res = append(res, '-')
	}
	n := numerator / denominator
	res = append(res, []byte(strconv.Itoa(n))...)
	if numerator%denominator == 0 {
		return string(res)
	}
	res = append(res, '.')
	foundIndex := make(map[int]int)
	for rest := numerator % denominator; rest > 0; rest %= denominator {
		if idx, ok := foundIndex[rest]; ok {
			res = append(res, 0)
			copy(res[idx+1:], res[idx:])
			res[idx] = '('
			res = append(res, ')')
			break
		}
		foundIndex[rest] = len(res)
		rest *= 10
		res = append(res, byte('0'+rest/denominator))
	}
	return string(res)
}
