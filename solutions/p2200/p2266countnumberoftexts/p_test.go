package p2266countnumberoftexts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countTexts(t *testing.T) {
	for _, tc := range []struct {
		pressedKeys string
		want        int
	}{
		{"22233", 8},
		{"222222222222222222222222222222222222", 82876089},
	} {
		t.Run(fmt.Sprintf("%+v", tc.pressedKeys), func(t *testing.T) {
			require.Equal(t, tc.want, countTexts(tc.pressedKeys))
		})
	}
}

var options = []int{
	2: 3,
	3: 3,
	4: 3,
	5: 3,
	6: 3,
	7: 4,
	8: 3,
	9: 4,
}

func countTexts(pressedKeys string) int {
	if len(pressedKeys) == 0 {
		return 0
	}
	const mod = 1e9 + 7
	res := 1
	n := len(pressedKeys)
	count := 1
	for i := 1; i <= n; i++ {
		if i < n && pressedKeys[i] == pressedKeys[i-1] {
			count++
			continue
		}

		nopt := options[pressedKeys[i-1]-'0']
		dp := make([]int, count+1)
		dp[0] = 1
		for j := 1; j <= count; j++ {
			for k := 1; k <= nopt; k++ {
				if j-k < 0 {
					break
				}
				dp[j] = (dp[j] + dp[j-k]) % mod
			}
		}
		res = (res * dp[count]) % mod

		count = 1
	}
	return res
}
