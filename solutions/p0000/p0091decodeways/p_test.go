package p0091decodeways

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numDecodings(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want int
	}{
		{"12", 2},
		{"226", 3},
		{"0", 0},
		{"1", 1},
		{"2101", 1},
	} {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, numDecodings(tc.in))
		})
	}
}

var oneTab = [10]byte{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}
var twoTab = [10][10]byte{
	1: {1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	2: {1, 1, 1, 1, 1, 1, 1},
}

func numDecodings(s string) int {
	n := len(s)
	var dp [3]int
	dp[1] = 1
	dp[2] = int(oneTab[s[0]-'0'])
	for x := 2; x <= n; x++ {
		dp[0], dp[1], dp[2] = dp[1], dp[2], 0
		dp[2] += int(oneTab[s[x-1]-'0']) * dp[1]
		dp[2] += int(twoTab[s[x-2]-'0'][s[x-1]-'0']) * dp[0]
	}
	return dp[2]
}
