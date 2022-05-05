package p1859sortingthesentence

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortSentence(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"is2 sentence4 This1 a3", "This is a sentence"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, sortSentence(tc.s))
		})
	}
}

func sortSentence(s string) string {
	parts := strings.Split(s, " ")
	sort.Slice(parts, func(i, j int) bool {
		return parts[i][len(parts[i])-1] < parts[j][len(parts[j])-1]
	})
	for i, p := range parts {
		parts[i] = p[:len(p)-1]
	}
	return strings.Join(parts, " ")
}
