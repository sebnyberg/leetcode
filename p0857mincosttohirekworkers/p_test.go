package p0857mincosttohirekworkers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mincostToHireWorkers(t *testing.T) {
	for _, tc := range []struct {
		quality []int
		wage    []int
		K       int
		want    float64
	}{
		// {[]int{10, 20, 5}, []int{70, 50, 30}, 3, 105},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.quality, tc.wage, tc.K), func(t *testing.T) {
			require.Equal(t, tc.want, mincostToHireWorkers(tc.quality, tc.wage, tc.K))
		})
	}
}
func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	return 0
}
