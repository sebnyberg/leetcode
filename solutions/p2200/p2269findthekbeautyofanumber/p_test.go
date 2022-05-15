package p2269findthekbeautyofanumber

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_divisorSubstrings(t *testing.T) {
	for _, tc := range []struct {
		num  int
		k    int
		want int
	}{
		{240, 2, 2},
		{430043, 2, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, divisorSubstrings(tc.num, tc.k))
		})
	}
}

func divisorSubstrings(num int, k int) int {
	s := fmt.Sprint(num)
	var res int
	for i := 0; i < len(s)-k+1; i++ {
		x, err := strconv.Atoi(s[i : i+k])
		if err != nil {
			panic(err)
		}
		if x == 0 {
			continue
		}
		if num%x == 0 {
			res++
		}
	}
	return res
}
