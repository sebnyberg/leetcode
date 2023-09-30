package p2744findmaximumnumberofstringpairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumNumberOfStringPairs(t *testing.T) {
	for i, tc := range []struct {
		words []string
		want  int
	}{
		{[]string{"cd", "ac", "dc", "ca", "zz"}, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maximumNumberOfStringPairs(tc.words))
		})
	}
}

func maximumNumberOfStringPairs(words []string) int {
	rev := make(map[string]int)
	reverse := func(s string) string {
		bs := []byte(s)
		for l, r := 0, len(bs)-1; l < r; l, r = l+1, r-1 {
			bs[l], bs[r] = bs[r], bs[l]
		}
		return string(bs)
	}

	var res int
	for _, w := range words {
		r := reverse(w)
		if rev[r] > 0 {
			res++
			rev[r]--
			continue
		}
		rev[w]++
	}
	return res
}
