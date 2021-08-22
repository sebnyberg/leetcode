package p1977numberofwaystoseparatenumber

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfCombinations(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want int
	}{
		{strings.Repeat("1", 3500), 755568658},
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
	n := len(num)

	lastNums := make([][]int, n+1)
	presums := make([][]int, n+1)
	lastNums[0] = append(lastNums[0], 0)
	presums[0] = []int{0, 1}

	for end := 1; end <= len(num); end++ {
		presums[end] = append(presums[end], 0)
		var presumIdx int
		for curStart := end - 1; curStart >= 0; curStart-- {
			if num[curStart] == '0' {
				continue
			}
			val := num[curStart:end]

			i := sort.Search(len(lastNums[curStart]), func(j int) bool {
				prevStart := lastNums[curStart][j]
				return gt(num[prevStart:curStart], val)
			})
			if count := presums[curStart][i]; count > 0 {
				lastNums[end] = append(lastNums[end], curStart)
				presums[end] = append(presums[end], presums[end][presumIdx]+count)
				presums[end][presumIdx] %= mod
				presumIdx++
			}
		}
	}
	return presums[n][len(presums[n])-1] % mod
}

func gt(a, b string) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		return a[i] > b[i]
	}
	return false
}
