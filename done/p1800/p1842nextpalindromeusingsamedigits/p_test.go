package p1842nextpalindromeusingsamedigits

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextPalindrome(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want string
	}{
		{"23143034132", "23314041332"},
		{"1221", "2112"},
		{"32123", ""},
		{"45544554", "54455445"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, nextPalindrome(tc.num))
		})
	}
}

func nextPalindrome(num string) string {
	n := len(num)
	mid := n / 2
	left := nextBigger(num[:mid])
	if left == "" {
		return ""
	}
	if n%2 == 0 {
		res := left + rev(left)
		return res
	}
	return left + string(num[mid]) + rev(left)
}

func nextBigger(num string) string {
	n := len(num)
	var i int
	for i = n - 2; i >= 0 && num[i] >= num[i+1]; i-- {
	}
	if i == -1 {
		return ""
	}
	// find smallest value greater than the one to swap
	firstGreater := byte('9' + 2)
	firstGreaterIdx := i + 1
	for j := i + 1; j < n; j++ {
		if num[j] > num[i] && num[j] < firstGreater {
			firstGreater = num[j]
			firstGreaterIdx = j
		}
	}
	bs := []byte(num)
	bs[i], bs[firstGreaterIdx] = bs[firstGreaterIdx], bs[i]
	sort.Slice(bs[i+1:], func(j, k int) bool {
		return bs[i+1+j] < bs[i+1+k]
	})
	return string(bs)
}

func rev(s string) string {
	bs := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	return string(bs)
}
