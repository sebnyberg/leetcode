package p2129capitalizethetitle

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_capitalizeTitle(t *testing.T) {
	for _, tc := range []struct {
		title string
		want  string
	}{
		{"capiTalIze tHe titLe", "Capitalize The Title"},
		{"First leTTeR of EACH Word", "First Letter of Each Word"},
		{"i lOve leetcode", "i Love Leetcode"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.title), func(t *testing.T) {
			require.Equal(t, tc.want, capitalizeTitle(tc.title))
		})
	}
}

func capitalizeTitle(title string) string {
	parts := strings.Fields(title)
	for i := range parts {
		parts[i] = strings.ToLower(parts[i])
		if len(parts[i]) > 2 {
			first := bytes.ToUpper([]byte{parts[i][0]})
			parts[i] = string(first) + parts[i][1:]
		}
	}
	return strings.Join(parts, " ")
}
