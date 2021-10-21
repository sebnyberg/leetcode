package p0848shiftingletters

import (
	"fmt"
	"testing"

	_ "github.com/bcmills/unsafeslice"
	"github.com/stretchr/testify/require"
)

func Test_shiftingLetters(t *testing.T) {
	for _, tc := range []struct {
		s      string
		shifts []int
		want   string
	}{
		{"a", []int{26}, "a"},
		{"abc", []int{3, 5, 9}, "rpl"},
		{"aaa", []int{1, 2, 3}, "gfd"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, shiftingLetters(tc.s, tc.shifts))
		})
	}
}

func shiftingLetters(s string, shifts []int) string {
	var val int
	n := len(shifts)
	bs := []byte(s)
	res := make([]byte, n)
	for i := len(shifts) - 1; i >= 0; i-- {
		val = (val + shifts[i]) % 26
		res[i] = 'a' + (bs[i]-'a'+byte(val))%26
	}
	return string(res)
}
