package p0132palindromepartitioning2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_partition(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"aab", 1},
		{"a", 0},
		{"ab", 1},
		{"apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp", 452},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, minCut(tc.s))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCut(s string) int {
	n := len(s)
	cut := make([]int, n+1) // # cuts for first k characters
	for i := 0; i <= n; i++ {
		cut[i] = i - 1
	}

	for i := 0; i < n; i++ {
		for j := 0; i-j >= 0 && i+j < n && s[i-j] == s[i+j]; j++ {
			cut[i+j+1] = min(cut[i+j+1], 1+cut[i-j])
		}
		for j := 1; i-j+1 >= 0 && i+j < n && s[i-j+1] == s[i+j]; j++ {
			cut[i+j+1] = min(cut[i+j+1], 1+cut[i-j+1])
		}
	}

	return cut[n]
}
