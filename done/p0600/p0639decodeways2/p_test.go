package p0639decodeways2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numDecodings(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"*********", 291868912},
		{"9617278761", 2},
		{"*0**0", 36},
		{"*1*1*0", 404},
		{"*1", 11},
		{"**", 96},
		{"*", 9},
		{"1*", 18},
		{"2*", 15},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, numDecodings(tc.s))
		})
	}
}

const mod = 1e9 + 7

func numDecodings(s string) int {
	n := len(s)
	bs := []byte(s)
	var dp [2]int
	dp[0], dp[1] = 1, 1
	for i := 0; i < n; i++ {
		ways := dp[1] * waysToDecodeOne(bs[i])
		if i > 0 {
			ways += dp[0] * waysToDecodeTwo(bs[i-1:i+1])
		}
		ways %= mod
		dp[0], dp[1] = dp[1], ways
	}
	return dp[1]
}

func waysToDecodeOne(b byte) int {
	switch b {
	case '0':
		return 0
	case '*':
		return 9
	default:
		return 1
	}
}

func waysToDecodeTwo(bs []byte) int {
	if len(bs) != 2 {
		panic("must be length 2")
	}
	a, b := bs[0], bs[1]
	switch {
	case a == '1':
		if b == '*' {
			return 9
		}
		return 1
	case a == '2':
		switch {
		case b == '*':
			return 6
		case b >= '0' && b <= '6':
			return 1
		}
	case a == '*':
		switch {
		case b == '*':
			return 15
		case b <= '6' && b >= '0':
			return 2
		default: // implicit b >= '7'
			return 1
		}
	}
	return 0
}
