package p0038countandsay

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countAndSay(t *testing.T) {
	for _, tc := range []struct {
		in   int
		want string
	}{
		{1, "1"},
		{4, "1211"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, countAndSay(tc.in))
		})
	}
}

func countAndSay(n int) string {
	res := []byte("1")
	next := []byte{}
	for i := 1; i < n; i++ {
		next = next[:0]
		count := 0
		for j := range res {
			if j == len(res)-1 || res[j] != res[j+1] {
				next = append(next, fmt.Sprint(count+1)...)
				next = append(next, res[j])
				count = 0
			} else {
				count++
			}
		}
		res, next = next, res
	}
	return string(res)
}
