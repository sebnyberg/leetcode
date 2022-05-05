package p1987numberofuniquegoodsubsequences

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfUniqueGoodSubsequences(t *testing.T) {
	for _, tc := range []struct {
		binary string
		want   int
	}{
		{"001", 2},
		{"11", 2},
		{"101", 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.binary), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfUniqueGoodSubsequences(tc.binary))
		})
	}
}

const mod = 1_000_000_007

func numberOfUniqueGoodSubsequences(binary string) int {
	if binary == "0" {
		return 1
	}
	// For each character in binary
	// If the character is a zero, the total number of new combinations is equal
	// to the number of new combinations for the previous zero plus the total
	// number of new combinations since the previous zero.

	// To start off, go through the binary until there is a one.
	var i int
	stack := []binCount{{'0', 0, 0}, {'1', 1, 1}}
	res := 0
	for i = 0; i < len(binary); i++ {
		if binary[i] == '1' {
			res++
			i++
			break
		}
	}
	stackIdx := 1
	for ; i < len(binary); i++ {
		ch := binary[i]
		if ch != stack[stackIdx].ch {
			stack = append(stack, binCount{ch, 0,
				stack[stackIdx].count + stack[stackIdx-1].delta,
			})
			stackIdx++
			stack[stackIdx].delta %= mod
		}
		stack[stackIdx].count += stack[stackIdx].delta
		stack[stackIdx].count %= mod
		res += stack[stackIdx].delta
		res %= mod
	}
	if strings.Contains(binary, "0") {
		res++
	}
	return res % mod
}

type binCount struct {
	ch    byte
	count int
	delta int
}
