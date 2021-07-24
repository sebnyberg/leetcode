package p0639decodeways2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numDecodings(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"*", 9},
		{"1*", 18},
		{"2*", 15},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, numDecodings(tc.s))
		})
	}
}

func numDecodings(s string) int {

}

func ways(bs []byte) int {
	if bs[0] == '0' {
		return 0
	}
	switch len(bs) {
	case 1:
		if bs[0] == '*' {
			return 9
		} else {
			return 1
		}
	case 2:
		// '*x'
		if bs[0] == '*' {
			switch {
			// '**' => 26
			case bs[1] == '*':
				return 26
			}
			if bs[1] == '*' {
				return 26
			} 
				if bs[1] <= '6' {
					return 2
				} else {
					return 1
				}
			}
		}
		// bs[0] is not '*'
		switch {
		case bs[0] == '0' || bs[0] >= '3':
			return 0
		case bs[1] == '*':
			return 9
		// bs[0] is either '1' or '2'
		case bs[1] <= '2':
			return 9
		default:
		}
	default:
		panic("bs must be <= 2 in length")
	}
}
