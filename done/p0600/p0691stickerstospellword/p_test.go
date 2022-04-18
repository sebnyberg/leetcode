package p0691stickerstospellword

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minStickers(t *testing.T) {
	for _, tc := range []struct {
		stickers []string
		target   string
		want     int
	}{
		{[]string{"with", "example", "science"}, "thehat", 3},
		{[]string{"notice", "possible"}, "basicbasic", -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.stickers), func(t *testing.T) {
			require.Equal(t, tc.want, minStickers(tc.stickers, tc.target))
		})
	}
}

func minStickers(stickers []string, target string) int {
	// Seems like a DP-type problem.
	// Each sticker contains a count of certain characters
	// We want to combine stickers so that the sum of characters from all stickers
	// becomes equal to or greater than the sum of characters in target.

	// There is also the BFS-style approach, but it's likely to bee too memory
	// heavy.

	// The only time the target cannot be reached is when it contains a character
	// for which there exists no sticker.
	var characterStickers [26][][26]int

	for _, st := range stickers {
		var charCount [26]int
		for _, ch := range st {
			charCount[ch-'a']++
		}
		for i, c := range charCount {
			if c > 0 {
				characterStickers[i] = append(characterStickers[i], charCount)
			}
		}
	}
	var remains [26]int
	for _, ch := range target {
		remains[ch-'a']++
		if len(characterStickers[ch-'a']) == 0 {
			return -1
		}
	}

	mem := make(map[[26]int]int)
	res := dfs(mem, characterStickers, remains)
	return res
}

var done [26]int

func dfs(mem map[[26]int]int, characterStickers [26][][26]int, remains [26]int) int {
	if v, exists := mem[remains]; exists {
		return v
	}
	if remains == done {
		return 0
	}
	res := math.MaxInt32
	for ch, v := range remains {
		if v == 0 {
			continue
		}
		// Try each sticker which has this character in it
		for _, sticker := range characterStickers[ch] {
			cpy := remains
			for ch2, count := range sticker {
				cpy[ch2] = max(0, cpy[ch2]-count)
			}
			res = min(res, 1+dfs(mem, characterStickers, cpy))
		}
	}
	mem[remains] = res
	return mem[remains]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
