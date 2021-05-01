package p1844replacealldigitswithcharacters

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_replaceDigits(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"a1c1e1", "abcdef"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, replaceDigits(tc.s))
		})
	}
}

func replaceDigits(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			sb.WriteByte(s[i])
			continue
		}
		sb.WriteByte(s[i-1] + s[i] - '0')
	}
	return sb.String()
}
