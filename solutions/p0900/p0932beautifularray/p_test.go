package p0932beautifularray

import (
	"fmt"
	"testing"
)

func Test_beautifulArray(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want []int
	}{
		{4, []int{2, 1, 4, 3}},
		{5, []int{3, 1, 2, 5, 4}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			res := beautifulArray(tc.n)
			for k := 1; k < len(res)-1; k++ {
				for i := 0; i < k; i++ {
					for j := k + 1; j < len(res); j++ {
						if res[i]+res[j] == res[k]*2 {
							t.FailNow()
						}
					}
				}
			}
		})
	}
}

func beautifulArray(n int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{1}
	}
	left := beautifulArray(n/2 + n&1)
	right := beautifulArray(n / 2)
	res := make([]int, len(left)+len(right))
	for i := range left {
		res[i] = left[i]*2 - 1
	}
	for i := range right {
		j := len(left) + i
		res[j] = right[i] * 2
	}
	return res
}
