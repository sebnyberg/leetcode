package p2224minimumnumberofoperationstoconverttime

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_convertTime(t *testing.T) {
	for _, tc := range []struct {
		current string
		correct string
		want    int
	}{
		{"02:30", "04:35", 3},
		{"11:00", "11:01", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.current), func(t *testing.T) {
			require.Equal(t, tc.want, convertTime(tc.current, tc.correct))
		})
	}
}

func convertTime(current string, correct string) int {
	var h1, m1, h2, m2 int
	fmt.Sscanf(current, "%02d:%02d", &h1, &m1)
	fmt.Sscanf(correct, "%02d:%02d", &h2, &m2)
	m1 += h1 * 60
	m2 += h2 * 60
	if m1 > m2 {
		m2 += 24 * 60
	}
	d := m2 - m1
	res := d / 60
	d %= 60
	res += d / 15
	d %= 15
	res += d / 5
	d %= 5
	res += d
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
