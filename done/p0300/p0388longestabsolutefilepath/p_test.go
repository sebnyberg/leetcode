package p0388longestabsolutefilepath

import (
	"bytes"
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func Test_lengthLongestPath(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  int
	}{
		{"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext", 20},
		{"a", 0},
		{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", 32},
		{"file1.txt\nfile2.txt\nlongfile.txt", 12},
	} {
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			require.Equal(t, tc.want, lengthLongestPath(tc.input))
		})
	}
}

func lengthLongestPath(input string) int {
	bs := *(*[]byte)(unsafe.Pointer(&input))
	stack := make([]int, 0)
	var maxLen int
	for _, line := range bytes.Split(bs, []byte{'\n'}) {
		var ndir int
		for ndir = 0; line[ndir] == '\t'; ndir++ {
		}
		lineLen := len(line) - ndir
		if len(stack) <= ndir {
			stack = append(stack, lineLen)
		} else {
			stack[ndir] = lineLen
		}
		if bytes.Contains(line, []byte{'.'}) {
			count := ndir
			for i := 0; i <= ndir; i++ {
				count += stack[i]
			}
			if count > maxLen {
				maxLen = count
			}
		}
	}
	return maxLen
}
