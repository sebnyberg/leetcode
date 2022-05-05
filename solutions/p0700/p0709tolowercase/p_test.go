package p0709tolowercase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_toLowerCase(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"Hello", "hello"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, toLowerCase(tc.s))
		})
	}
}

const upperDist = 'a' - 'A'

func toLowerCase(s string) string {
	res := make([]rune, len(s))
	for i, ch := range s {
		if ch <= 'Z' && ch >= 'A' {
			ch += upperDist
		}
		res[i] = ch
	}
	return string(res)
}
