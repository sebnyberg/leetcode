package p0649dota2senate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_predictPartyVictory(t *testing.T) {
	for _, tc := range []struct {
		senate string
		want   string
	}{
		{"RD", "Radiant"},
		{"RDD", "Dire"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.senate), func(t *testing.T) {
			require.Equal(t, tc.want, predictPartyVictory(tc.senate))
		})
	}
}

func predictPartyVictory(senate string) string {
	// I assume this greedy solution will work:
	// Each senator bans the next-in-line senator from the other team.
	var toBan [2]int
	n := len(senate)
	banned := make([]bool, n)
	seen := [2]bool{true, true}
	const R, D = 0, 1
	v := [255]uint8{'R': 0, 'D': 1}
	for seen[R] && seen[D] {
		seen[R] = false
		seen[D] = false
		for i, ch := range senate {
			if banned[i] {
				continue
			}
			if toBan[v[ch]] > 0 {
				banned[i] = true
				toBan[v[ch]]--
				continue
			}
			toBan[1-v[ch]]++ // ban next, other player
			seen[v[ch]] = true
		}
	}
	if seen[R] {
		return "Radiant"
	}
	return "Dire"
}
