package p1977numberofwaystoseparatenumber

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfCombinations(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want int
	}{
		{strings.Repeat("1", 3500), 10},
		// {"24896", 6},
		// {"1203", 2},
		// {"327", 2},
		// {"094", 0},
		// {"0", 0},
		// {"9999999999999", 101},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfCombinations(tc.num))
		})
	}
}

const mod = 1e9 + 7

func numberOfCombinations(num string) int {
	// series := make([][]numSeries, len(num))
	// series[0] = append(series[0], numSeries{"0", 1, 0})
	series := []numSeries{{"0", 1, 0}}
	for i := 1; i <= len(num); i++ {
		newSeries := map[string]int{}
		for _, s := range series {
			if geq(num[s.pos:i], s.val) && num[s.pos] != '0' {
				newSeries[num[s.pos:i]] += s.count
				newSeries[num[s.pos:i]] %= mod
			}
		}
		for val, count := range newSeries {
			series = append(series, numSeries{val, count, i})
		}
	}
	var res int
	for _, s := range series {
		if s.pos == len(num) {
			res += s.count
			res %= mod
		}
	}
	return res % mod
}

func geq(a, b string) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		return a[i] > b[i]
	}
	return true
}

type numSeries struct {
	val   string
	count int
	pos   int
}
