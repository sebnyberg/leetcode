package p2021brightestpositiononstreet

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_brightestPosition(t *testing.T) {
	for _, tc := range []struct {
		lights [][]int
		want   int
	}{
		{[][]int{{-3, 2}, {1, 2}, {3, 3}}, -1},
		{[][]int{{1, 0}, {0, 1}}, 1},
		{[][]int{{1, 2}}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.lights), func(t *testing.T) {
			require.Equal(t, tc.want, brightestPosition(tc.lights))
		})
	}
}

func brightestPosition(lights [][]int) int {
	type lightDelta struct {
		position int
		diff     int8
	}
	deltas := make([]lightDelta, 0, len(lights)*2)
	for _, light := range lights {
		pos, lightRange := light[0], light[1]
		deltas = append(deltas, lightDelta{pos - lightRange, 1})
		deltas = append(deltas, lightDelta{pos + lightRange + 1, -1})
	}
	sort.Slice(deltas, func(i, j int) bool {
		return deltas[i].position < deltas[j].position
	})
	var brightness, maxBrightness, maxBrightnessPos int
	for i := 0; i < len(deltas); {
		j := i
		for j < len(deltas) && deltas[j].position == deltas[i].position {
			brightness += int(deltas[j].diff)
			j++
		}
		if brightness > maxBrightness {
			maxBrightnessPos = deltas[i].position
			maxBrightness = brightness
		}
		i = j
	}
	return maxBrightnessPos
}
