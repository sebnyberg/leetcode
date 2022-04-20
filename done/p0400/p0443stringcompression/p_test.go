package p0443stringcompression

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_compress(t *testing.T) {
	for _, tc := range []struct {
		chars []byte
		want  int
	}{
		{[]byte("abbbbbbbbbbbb"), 4},
		{[]byte("aabbccc"), 6},
		{[]byte("a"), 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.chars), func(t *testing.T) {
			require.Equal(t, tc.want, compress(tc.chars))
		})
	}
}

func compress(chars []byte) int {
	var j int
	count := 1
	write := func(ch byte) {
		chars[j] = ch
		j++
		if count > 1 {
			s := []byte(fmt.Sprint(count))
			j += copy(chars[j:], s)
		}
		count = 1
	}

	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			count++
			continue
		}
		write(chars[i-1])
	}
	write(chars[len(chars)-1])
	return j
}
