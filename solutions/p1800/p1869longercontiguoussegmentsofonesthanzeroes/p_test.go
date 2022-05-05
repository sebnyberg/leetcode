package p1869longercontiguoussegmentsofonesthanzeroes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkZeroOnes(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"1101", true},
		{"111000", false},
		{"110100010", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, checkZeroOnes(tc.s))
		})
	}
}

func checkZeroOnes(s string) bool {
	var curCount [2]int
	var maxCount [2]int
	for i := 0; i < len(s); i++ {
		a := s[i] - '0'
		if i == 0 {
			curCount[a]++
		} else {
			if s[i-1] != s[i] {
				curCount[s[i-1]-'0'] = 0
			}
			curCount[a]++
		}
		maxCount[a] = max(maxCount[a], curCount[a])
	}
	return maxCount[1] > maxCount[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
