package p1904thenumberoffullroundsyouhaveplayed

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfRounds(t *testing.T) {
	for _, tc := range []struct {
		startTime  string
		finishTime string
		want       int
	}{
		{"12:01", "12:44", 1},
		{"20:00", "06:00", 40},
		{"00:00", "23:59", 95},
	} {
		t.Run(fmt.Sprintf("%+v", tc.startTime), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfRounds(tc.startTime, tc.finishTime))
		})
	}
}

func numberOfRounds(startTime string, finishTime string) int {
	startH, _ := strconv.Atoi(startTime[0:2])
	startM, _ := strconv.Atoi(startTime[3:])
	startM += startH * 60

	finishH, _ := strconv.Atoi(finishTime[0:2])
	finishM, _ := strconv.Atoi(finishTime[3:])
	finishM += finishH * 60

	if finishM < startM {
		finishM += 24 * 60
	}

	// Round start minutes to next quarter
	if startM%15 != 0 {
		startM -= startM % 15
		startM += 15
	}

	d := finishM - startM
	return d / 15
}
