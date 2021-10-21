package p186reversewordsinastring2

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseWords(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			bs := []byte(tc.s)
			reverseWords(bs)
			require.Equal(t, tc.want, string(bs))
		})
	}
}

func reverseWords(s []byte) {
	ss := string(s)
	parts := strings.Split(ss, " ")
	n := len(parts)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	sss := strings.Join(parts, " ")
	for i := range s {
		s[i] = sss[i]
	}
}
