package p3494findtheminimumamountoftimetobrewpotions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minTime(t *testing.T) {
	for _, tc := range []struct {
		skill []int
		mana  []int
		want  int64
	}{
		{[]int{1, 5, 2, 4}, []int{5, 1, 4, 2}, 110},
	} {
		t.Run(fmt.Sprintf("%+v", tc.skill), func(t *testing.T) {
			require.Equal(t, tc.want, minTime(tc.skill, tc.mana))
		})
	}
}

func minTime(skill []int, mana []int) int64 {
	n := len(skill)
	times := make([]int, n)
	for i := range skill {
		times[i] = skill[i] * mana[0]
		if i > 0 {
			times[i] += times[i-1]
		}
	}
	for _, potion := range mana[1:] {
		var shift int
		t := times[0]
		for j := range skill {
			t += potion * skill[j]
			if j < len(times)-1 {
				shift = max(shift, times[j+1]-t)
			}
		}
		times[0] += potion*skill[0] + shift
		for i := 1; i < len(times); i++ {
			times[i] = times[i-1] + potion*skill[i]
		}
	}
	return int64(times[n-1])
}
