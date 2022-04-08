package p2214minimumhealthtobeatgame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumHealth(t *testing.T) {
	for _, tc := range []struct {
		damage []int
		armor  int
		want   int64
	}{
		{[]int{2, 7, 4, 3}, 4, 13},
		{[]int{2, 5, 3, 4}, 7, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.damage), func(t *testing.T) {
			require.Equal(t, tc.want, minimumHealth(tc.damage, tc.armor))
		})
	}
}

func minimumHealth(damage []int, armor int) int64 {
	var hp, hpNoArmor int
	for _, d := range damage {
		hpNoArmor = max(hpNoArmor-d, hp-max(0, d-armor))
		hp -= d
	}
	return int64(-hpNoArmor) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
