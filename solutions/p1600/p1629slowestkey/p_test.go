package p1629slowestkey

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_slowestKey(t *testing.T) {
	for _, tc := range []struct {
		releaseTimes []int
		keysPressed  string
		want         byte
	}{
		{[]int{9, 29, 49, 50}, "cbcd", 'c'},
		{[]int{12, 23, 36, 46, 62}, "spuda", 'a'},
	} {
		t.Run(fmt.Sprintf("%+v", tc.releaseTimes), func(t *testing.T) {
			require.Equal(t, tc.want, slowestKey(tc.releaseTimes, tc.keysPressed))
		})
	}
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	longestPress := releaseTimes[0]
	longestPressKey := keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		d := releaseTimes[i] - releaseTimes[i-1]
		if d > longestPress {
			longestPress = d
			longestPressKey = keysPressed[i]
		} else if d == longestPress {
			if longestPressKey < keysPressed[i] {
				longestPressKey = keysPressed[i]
			}
		}
	}
	return longestPressKey
}
