package p0808soupservings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_soupServings(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want float64
	}{
		{10000, 0.9999999},
		{1000, 0.976565},
		{50, 0.625},
		{100, 0.71875},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.InEpsilon(t, tc.want, soupServings(tc.n), 1e-5)
		})
	}
}

func soupServings(n int) float64 {
	if n >= 10000 {
		return 1
	}
	mem := make(map[[2]int]float64)
	res := dp(mem, n, n)
	return res
}

func dp(mem map[[2]int]float64, soupA, soupB int) float64 {
	k := [2]int{soupA, soupB}
	if v, exists := mem[k]; exists {
		return v
	}
	if soupA <= 0 || soupB <= 0 {
		if soupA <= 0 {
			if soupB > 0 {
				return 1
			}
			return 0.5
		}
		return 0
	}
	var res float64
	res += 0.25 * dp(mem, soupA-100, soupB)
	res += 0.25 * dp(mem, soupA-75, soupB-25)
	res += 0.25 * dp(mem, soupA-50, soupB-50)
	res += 0.25 * dp(mem, soupA-25, soupB-75)
	mem[k] = res
	return res
}
