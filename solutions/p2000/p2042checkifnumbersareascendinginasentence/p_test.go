package p2042checkifnumbersareascendinginasentence

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_areNumbersAscending(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"hello world 5 x 5", false},
		{"1 box has 3 blue 4 red 6 green and 12 yellow marbles", true},
		{"sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s", false},
		{"4 5 11 26", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, areNumbersAscending(tc.s))
		})
	}
}

func areNumbersAscending(s string) bool {
	parts := strings.Split(s, " ")
	prev := -1
	for _, part := range parts {
		if n, err := strconv.Atoi(part); err == nil {
			if n <= prev {
				return false
			}
			prev = n
		}
	}
	return true
}
