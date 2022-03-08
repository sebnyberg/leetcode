package p0609findduplicatefileinsystem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDuplicate(t *testing.T) {
	for _, tc := range []struct {
		paths []string
		want  [][]string
	}{
		{
			[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
			[][]string{{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}, {"root/a/1.txt", "root/c/3.txt"}},
		},
		{
			[]string{"root/a 1.txt(abcd) 2.txt(efsfgh)", "root/c 3.txt(abdfcd)", "root/c/d 4.txt(efggdfh)"},
			[][]string{},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.paths), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, findDuplicate(tc.paths))
		})
	}
}

func findDuplicate(paths []string) [][]string {
	contentFiles := make(map[string][]string)
	for _, p := range paths {
		parts := strings.Fields(p)
		prefix := parts[0]
		for _, f := range parts[1:] {
			pp := strings.Split(f, "(")
			fname := pp[0]
			contents := pp[1][:len(pp[1])-1]
			contentFiles[contents] = append(contentFiles[contents], prefix+"/"+fname)
		}
	}
	var res [][]string
	for _, files := range contentFiles {
		if len(files) > 1 {
			res = append(res, files)
		}
	}
	return res
}
