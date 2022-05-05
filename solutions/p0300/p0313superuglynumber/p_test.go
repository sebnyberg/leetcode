package p0313superuglynumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nthSuperUglyNumber(t *testing.T) {
	for _, tc := range []struct {
		n      int
		primes []int
		want   int
	}{
		{12, []int{2, 7, 13, 19}, 32},
		{1, []int{2, 3, 5}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, nthSuperUglyNumber(tc.n, tc.primes))
		})
	}
}

func nthSuperUglyNumber(n int, primes []int) int {
	nums := make([]uint32, n)
	pprimes := make([]uint32, len(primes))
	for i, prime := range primes {
		pprimes[i] = uint32(prime)
	}

	offsets := make([]uint32, len(primes))
	nums[0] = 1
	for i := 1; i < n; i++ {
		for j, p := range pprimes {
			offset := offsets[j]
			v := nums[offset] * p
			if nums[i-1] == v { // duplicate
				offset++
				offsets[j] = offset
				v = nums[offset] * p
			}

			if nums[i] == 0 || v <= nums[i] {
				nums[i] = v
			}
		}
	}
	return int(nums[n-1])
}
