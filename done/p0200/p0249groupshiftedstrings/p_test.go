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
			require.Equal(t, tc.want, groupStrings(tc.strings))
		})
	}
}

func groupStrings(strings []string) [][]string {

}
