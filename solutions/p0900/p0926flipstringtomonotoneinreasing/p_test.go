package p0926flipstringtomonotoneinreasing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minFlipsMonoIncr(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"00110", 1},
		{"010110", 2},
		{"00011000", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, minFlipsMonoIncr(tc.s))
		})
	}
}

func minFlipsMonoIncr(s string) int {
	// Do RLE of the string
	rle := encodingStack{{0, 123}} // sentinel
	for i := range s {
		ch := s[i]
		if ch != rle.peek().val {
			rle.push(&lengthEncoding{1, ch})
		} else {
			rle.peek().width++
		}
	}

	// Calculate the cost of making the string mono incrementing by keeping track
	// of the cost of having the remainder be all ones, or the remainder being
	// zeroes then possibly ones.
	var costAllOnes, costMonoIncr int
	for len(rle) > 1 {
		x := rle.pop()
		if x.val == '0' {
			// mono incr cost remains the same
			costAllOnes += x.width
		} else {
			costMonoIncr = min(
				costMonoIncr+x.width,
				costAllOnes,
			)
		}
	}

	return costMonoIncr
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type lengthEncoding struct {
	width int
	val   byte
}

type encodingStack []*lengthEncoding

func (s encodingStack) peek() *lengthEncoding {
	return s[len(s)-1]
}

func (s *encodingStack) push(e *lengthEncoding) {
	*s = append(*s, e)
}

func (s *encodingStack) pop() *lengthEncoding {
	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return el
}
