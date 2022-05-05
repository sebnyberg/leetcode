package p0609findduplicateinsystem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDuplicateInSystem(t *testing.T) {
	for _, tc := range []struct {
		paths []string
		want  [][]string
	}{
		{
			[]string{
				"root/a 1.txt(abcd) 2.txt(efgh)",
				"root/c 3.txt(abcd)",
				"root/c/d 4.txt(efgh)",
				"root 4.txt(efgh)",
			},
			[][]string{
				{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"},
				{"root/a/1.txt", "root/c/3.txt"},
			},
		},
		{
			[]string{
				"root/a 1.txt(abcd) 2.txt(efgh)",
				"root/c 3.txt(abcd)",
				"root/c/d 4.txt(efgh)",
			},
			[][]string{
				{"root/a/2.txt", "root/c/d/4.txt"},
				{"root/a/1.txt", "root/c/3.txt"},
			},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.paths), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, findDuplicate(tc.paths))
		})
	}
}

func findDuplicate(paths []string) [][]string {
	contentFiles := make(map[string][]string)
	for _, path := range paths {
		parts := strings.Split(path, " ")
		dir := parts[0]
		for _, fpath := range parts[1:] {
			fparts := strings.Split(fpath, "(")
			filename := fparts[0]
			content := fparts[1][:len(fparts[1])-1]
			contentFiles[content] = append(contentFiles[content], dir+"/"+filename)
		}
	}
	res := make([][]string, 0)
	for _, files := range contentFiles {
		if len(files) <= 1 {
			continue
		}
		res = append(res, files)
	}
	return res
}
