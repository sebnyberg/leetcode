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
		{"24896", 6},
		{"1203", 2},
		{"327", 2},
		{"094", 0},
		{"0", 0},
		{"9999999999999", 101},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfCombinations(tc.num))
		})
	}
}

const mod = 1e9 + 7

func numberOfCombinations(num string) int {
	// Keep list of "last numbers" from previous series in a list
	// The index in the list denotes where the series ended.
	// Each list is sorted by start index so that binary search can rule out
	// exactly how many previous series are relevant for a given position in the
	// nums string.
	series := make([][]lastNum, len(num)+1)
	presums := make([][]int, len(num)+1)
	series[0] = append(series[0], lastNum{0, 0, 1})
	presums[0] = make([]int, 2)
	presums[0][1] = 1
	for i := 1; i <= len(num); i++ {
		// For each prior end position, count the number of series which end in
		// a number which is smaller than or equal to the current value
		for j := 0; j < i; j++ {
			if num[j] == '0' {
				continue
			}
			var count int
			val := num[j:i]
			for _, s := range series[j] {
				if geq(val, num[s.start:j]) {
					count += s.count
				}
			}
			if count > 0 {
				series[i] = append(series[i], lastNum{j, i, count})
			}
		}
	}
	var res int
	for _, s := range series[len(num)] {
		res += s.count
	}
	return res
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

// lastNum stores the last number of a series of numbers, and a count of the
// number of ways one could end up with such a series.
type lastNum struct {
	start, end int
	count      int
}
