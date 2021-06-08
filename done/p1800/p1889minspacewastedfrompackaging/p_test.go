package p1889minspacewastedfrompackaging

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minWastedSpace(t *testing.T) {
	for _, tc := range []struct {
		packages []int
		boxes    [][]int
		want     int
	}{
		{[]int{2, 3, 5}, [][]int{{4, 8}, {2, 8}}, 6},
		{[]int{2, 3, 5}, [][]int{{1, 4}, {2, 3}, {3, 4}}, -1},
		{[]int{3, 5, 8, 10, 11, 12}, [][]int{{12}, {11, 9}, {10, 5, 14}}, 9},
		// {[]int{3, 5, 8, 10, 11, 12}, [][]int{{1, 5, 6, 7, 9, 9, 10, 15, 16, 22}}, 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.packages), func(t *testing.T) {
			require.Equal(t, tc.want, minWastedSpace(tc.packages, tc.boxes))
		})
	}
}

const mod = 1000000007

func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	n := len(packages)
	var maxWasted int
	for _, pkg := range packages {
		maxWasted += pkg
	}

	minBoxUsed := math.MaxInt64
	for _, b := range boxes {
		sort.Ints(b)
		if b[len(b)-1] < packages[n-1] {
			continue
		}
		var curBoxUsed, offset int
		for _, box := range b {
			pos := binSearch(packages, box+1)
			// Add all packages from this position
			curBoxUsed += box * (pos - offset)
			offset = pos
		}
		if curBoxUsed < minBoxUsed {
			minBoxUsed = curBoxUsed
		}
	}
	if minBoxUsed < math.MaxInt64 {
		return (minBoxUsed - maxWasted) % mod
	}
	return -1
}

func binSearch(a []int, b int) int {
	l, r := 0, len(a)
	for l < r {
		m := (l + r) / 2
		if a[m] < b {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
