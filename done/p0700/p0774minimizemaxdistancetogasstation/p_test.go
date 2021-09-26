package p0774minimizemaxdistancetogasstation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minmaxGasDist(t *testing.T) {
	for _, tc := range []struct {
		stations []int
		k        int
		want     float64
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 0.5},
		{[]int{23, 24, 36, 39, 46, 56, 57, 65, 84, 98}, 1, 14},
	} {
		t.Run(fmt.Sprintf("%+v", tc.stations), func(t *testing.T) {
			require.InEpsilon(t, tc.want, minmaxGasDist(tc.stations, tc.k), 1e-6)
		})
	}
}

const eps = 1e-6

func minmaxGasDist(stations []int, k int) float64 {
	// This is a typical binary search exercise. Set a low and a high value. Take
	// a mid point and check whether the constraints are satisfiable given that
	// solution. If it is true, then mid becomes the new high, otherwise mid
	// becomes the new low. Once the distance between low and high is smaller than
	// the error margin (epsilon), exit.
	n := len(stations)
	lo, hi := 0.0, float64(stations[n-1]) // not best choice of high but less code
	for hi-lo > eps {
		mid := (hi + lo) / 2
		m := 0
		for i := 1; i < n; i++ {
			m += int(float64(stations[i]-stations[i-1]) / mid)
			if m > k {
				break
			}
		}
		if m <= k {
			hi = mid
		} else {
			lo = mid
		}
	}
	return (hi + lo) / 2
}
