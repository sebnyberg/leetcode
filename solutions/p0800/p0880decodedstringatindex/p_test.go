package p0880decodedstringatindex

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_decodeAtIndex(t *testing.T) {
	for i, tc := range []struct {
		s    string
		k    int
		want string
	}{
		{"a2345678999999999999999", 1, "a"},
		{"ha22", 5, "h"},
		{"leet2code3", 10, "o"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, decodeAtIndex(tc.s, tc.k))
		})
	}
}

func decodeAtIndex(s string, k int) string {
	n := []int{0}
	ss := [][]byte{{}}
	repeat := []int{1}
	var j int
	for i := range s {
		if unicode.IsDigit(rune(s[i])) {
			repeat[j] *= int(s[i] - '0')
			continue
		}
		if repeat[j] != 1 {
			j++
			repeat = append(repeat, 1)
			ss = append(ss, []byte{})
			n = append(n, 0)
		}
		ss[j] = append(ss[j], s[i])
		n[j]++
	}
	totLen := make([]int, len(n))
	totLen[0] = n[0]
	for i := 1; i < len(totLen); i++ {
		totLen[i] = n[i] + totLen[i-1]*repeat[i-1]
	}
	k--
	for j := len(totLen) - 1; j >= 0; j-- {
		v := totLen[j]
		k %= v
		left := v - n[j]
		if k < left {
			continue
		}
		res := ss[j][k-left]
		return string(res)
	}
	return ""
}
