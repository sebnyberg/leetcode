package p3186maximumtotaldamagewithspellcasting

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumTotalDamage(t *testing.T) {
	for _, tc := range []struct {
		power []int
		want  int64
	}{
		{[]int{1, 1, 3, 4}, 6},
		{[]int{7, 1, 6, 6}, 13},
	} {
		t.Run(fmt.Sprintf("%+v", tc.power), func(t *testing.T) {
			require.Equal(t, tc.want, maximumTotalDamage(tc.power))
		})
	}
}

func maximumTotalDamage(power []int) int64 {
	// This is a typical DP problem, and it can be done with bottom-up, but it's
	// easier to do top-down.

	sort.Ints(power) // Sort in ascending order

	n := len(power)
	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}
	return int64(dp(mem, power, 0))
}

func dp(mem, power []int, i int) int {
	if i == len(power) {
		return 0
	}
	if mem[i] != -1 {
		return mem[i]
	}
	// Either cast or don't cast this spell

	// If we don't cast this spell, then we must find the next spell that is not
	// of the same power as this one
	j := i
	for ; j < len(power) && power[i] == power[j]; j++ {
	}
	noCast := dp(mem, power, j)

	// If we cast the spell, then we must skip spells until power[j]-2 > power[i]
	cast := power[i] * (j - i)
	for ; j < len(power) && power[j]-2 <= power[i]; j++ {
	}
	cast += dp(mem, power, j)
	res := max(cast, noCast)
	mem[i] = res
	return mem[i]
}
