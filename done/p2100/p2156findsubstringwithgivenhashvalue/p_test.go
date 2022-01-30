package p2156findsubstringwithgivenhashvalue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subStrHash(t *testing.T) {
	for _, tc := range []struct {
		s         string
		power     int
		modulo    int
		k         int
		hashValue int
		want      string
	}{
		{"xqgfatvtlwnnkxipmipcpqwbxihxblaplpfckvxtihonijhtezdnkjmmk", 22, 51, 41, 9, "xqgfatvtlwnnkxipmipcpqwbxihxblaplpfckvxti"},
		{"leetcode", 7, 20, 2, 0, "ee"},
		{"leetcode", 7, 20, 2, 0, "ee"},
		{"fbxzaad", 31, 100, 3, 32, "fbx"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, subStrHash(tc.s, tc.power, tc.modulo, tc.k, tc.hashValue))
		})
	}
}

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	// Going from right to left, it is easy to re-hash
	var h int
	n := len(s)
	rmPow := 1
	for i := 0; i < k; i++ {
		rmPow = (rmPow * power) % modulo
	}
	var res string
	for i := n - 1; i >= 0; i-- {
		// Add to hash
		h = (h*power + int(s[i]-'a'+1)) % modulo
		if n-i < k {
			continue
		}
		// Remove tail
		if n-i > k {
			h = (h - (int(s[i+k]-'a'+1)*rmPow)%modulo + modulo) % modulo
		}
		if h == hashValue {
			res = s[i : i+k]
		}
	}
	return res
}
