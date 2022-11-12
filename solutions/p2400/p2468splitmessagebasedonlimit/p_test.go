package p2468splitmessagebasedonlimit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_splitMessage(t *testing.T) {
	for i, tc := range []struct {
		message string
		limit   int
		want    []string
	}{
		{
			"this is really a very awesome message",
			9,
			[]string{"thi<1/14>", "s i<2/14>", "s r<3/14>", "eal<4/14>", "ly <5/14>", "a v<6/14>", "ery<7/14>", " aw<8/14>", "eso<9/14>", "me<10/14>", " m<11/14>", "es<12/14>", "sa<13/14>", "ge<14/14>"},
		},
		// {
		// 	"short message",
		// 	15,
		// 	[]string{"short mess<1/2>", "age<2/2>"},
		// },
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, splitMessage(tc.message, tc.limit))
		})
	}
}

func splitMessage(message string, limit int) []string {
	n := len(message)
	// Assume that the answer uses 1 digit, i.e. <x/x>,
	// Then we must cover n within the first 9 elements (at most)
	for m := 1; m <= 5; m++ {
		nn := n
		a := 10
		var count int
		for x := 1; x <= m; x++ {
			perPart := (limit - (3 + x + m))
			if perPart <= 0 {
				return []string{}
			}
			maxNParts := a - a/10
			if nn <= maxNParts*perPart {
				if x < m {
					goto cont
				}
				// found a match, construct message
				count += nn / perPart
				if nn%perPart != 0 {
					count++
				}
				var res []string
				var i int
				for c := 1; c <= count; c++ {
					aaa := limit - m - 3 - len(fmt.Sprint(c))
					j := i + aaa
					if j > n {
						j = n
					}
					res = append(res, fmt.Sprintf("%s<%d/%d>", message[i:j], c, count))
					i = j
				}
				return res
			}
			nn -= maxNParts * perPart
			count += maxNParts

			a *= 10
		}

	cont:
	}
	return []string{}
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	return pow(a, b/2) * pow(a, b/2+b&1)
}
