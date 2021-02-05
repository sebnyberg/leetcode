package p0071simplifypath

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_simplifyPath(t *testing.T) {
	for _, tc := range []struct {
		path string
		want string
	}{
		{"/home/", "/home"},
		{"/home///foo/", "/home/foo"},
		{"/home/./foo/", "/home/foo"},
		{"/home//./foo/", "/home/foo"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.path), func(t *testing.T) {
			require.Equal(t, tc.want, simplifyPath(tc.path))
		})
	}
}

func simplifyPath(path string) string {
	path = strings.TrimRight(path, "/")

	pathParts := strings.Split(path, "/")
	for i := 0; ; {
		if i == len(pathParts) {
			break
		}
		if pathParts[i] == "" || pathParts[i] == "." {
			pathParts = append(pathParts[:i], pathParts[i+1:]...)
			continue
		}
		if pathParts[i] == ".." {
			if i == 0 {
				pathParts = pathParts[1:]
				continue
			}
			if i == 1 {
				pathParts = pathParts[2:]
				i--
				continue
			}
			pathParts = append(pathParts[:i-1], pathParts[i+1:]...)
			i--
			continue
		}
		i++
	}

	return "/" + strings.Join(pathParts, "/")
}
