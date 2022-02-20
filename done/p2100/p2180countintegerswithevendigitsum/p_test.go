package p2180

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countEven(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want int
	}{
		{4, 2},
		{30, 14},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, countEven(tc.num))
		})
	}
}

func countEven(num int) int {
	var res int
	for x := 1; x <= num; x++ {
		var sum int
		s := fmt.Sprint(x)
		for _, ch := range s {
			sum += int(ch - '0')
		}
		if sum%2 == 0 {
			res++
		}
	}
	return res
}
