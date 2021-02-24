package p0132palindromepartitioning2

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_partition(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		// {"aab", 1},
		// {"a", 0},
		// {"ab", 1},
		// {"apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, minCut(tc.s))
		})
	}
}

func minCut(s string) int {
	var f partitionFinder
	f.partitions = math.MaxInt32
	return f.partitions - 1
}

type partitionFinder struct {
	partitions  int
	palindromes map[string]bool
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (f *partitionFinder) findPartitions(s string, partitions int) {
	if partitions >= f.partitions {
		return
	}
	if len(s) == 0 {
		f.partitions = min(f.partitions, partitions)
		return
	}
	for i := len(s) - 1; i > 0; i-- {
		if _, exists := f.palindromes[s[:i]]; !exists {
			f.palindromes[s[:i]] = isPalindrome(s[:i])
		}
		if f.palindromes[s[:i]] {
			f.findPartitions(s[i:], partitions+1)
		}
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < (len(s)+1)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
