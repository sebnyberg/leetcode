package p0249groupshiftedstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_groupStrings(t *testing.T) {
	for _, tc := range []struct {
		strings []string
		want    [][]string
	}{
		{[]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}, [][]string{{"acef"}, {"a", "z"}, {"abc", "bcd", "xyz"}, {"az", "ba"}}},
		{[]string{"a"}, [][]string{{"a"}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.strings), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, groupStrings(tc.strings))
		})
	}
}

func groupStrings(strings []string) [][]string {
	// Normalize each string so that the first letter is always "a"
	// Then collect strings into a joint map
	results := make(map[string][]string)
	for _, s := range strings {
		norm := normalize(s)
		results[norm] = append(results[norm], s)
	}
	resultsList := make([][]string, 0, len(results))
	for _, v := range results {
		resultsList = append(resultsList, v)
	}
	return resultsList
}

func normalize(s string) string {
	if s[0] == 'a' { // already normalized
		return s
	}
	// normalize based on 'a'
	res := make([]byte, len(s))
	d := s[0] - 'a'
	for i, ch := range s {
		resCh := byte(ch) - d
		if resCh < byte('a') {
			resCh += 26
		}
		res[i] = resCh
	}
	return string(res)
}
