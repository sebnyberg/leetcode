package p0443stringcompression

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_compress(t *testing.T) {
	for _, tc := range []struct {
		chars string
		want  int
	}{
		{"aabbccc", 6},
		{"a", 1},
		{"abbbbbbbbbbbb", 4},
		{"aaabbaa", 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.chars), func(t *testing.T) {
			require.Equal(t, tc.want, compress([]byte(tc.chars)))
		})
	}
}

func compress(chars []byte) int {
	curCh := chars[0]
	count := 1
	var idx int
	addRes := func(count int) {
		chars[idx] = curCh
		idx++
		if count > 1 {
			nStr := strconv.Itoa(count)
			for i := range nStr {
				chars[idx] = nStr[i]
				idx++
			}
		}
	}
	for i := range chars {
		if i == 0 {
			continue
		}
		if chars[i] == curCh {
			count++
			continue
		}
		addRes(count)
		curCh = chars[i]
		count = 1
	}
	addRes(count)
	chars = chars[:idx]
	return len(chars)
}
