package p0667beautifularrangement2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_constructArray(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want []int
	}{
		{3, 1, []int{1, 2, 3}},
		{3, 2, []int{1, 3, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, constructArray(tc.n, tc.k))
		})
	}
}

func constructArray(n int, k int) []int {
	res := []int{1}
	i := 1
	prev := 1
	for k > 0 {
		if i%2 == 1 { // odd number, increase
			res = append(res, prev+k)
			prev = prev + k
		} else {
			res = append(res, prev-k)
			prev = prev - k
		}
		k--
		i++
	}
	cur := len(res)
	for cur < n {
		cur++
		res = append(res, cur)
	}
	return res
}
