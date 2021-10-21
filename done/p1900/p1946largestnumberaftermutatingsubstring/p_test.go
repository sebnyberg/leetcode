package p1946largestnumberaftermutatingsubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumNumber(t *testing.T) {
	for _, tc := range []struct {
		num    string
		change []int
		want   string
	}{
		{"132", []int{9, 8, 5, 0, 3, 6, 4, 2, 6, 8}, "832"},
		{"021", []int{9, 4, 3, 5, 7, 2, 1, 9, 0, 6}, "934"},
		{"5", []int{1, 4, 7, 5, 3, 2, 5, 6, 9, 4}, "5"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, maximumNumber(tc.num, tc.change))
		})
	}
}

func maximumNumber(num string, change []int) string {
	// Find first place in num where a change would increase the value, this is
	// the optimal change. Then, change values for as long as the change improves
	// the solution.
	res := make([]byte, 0, len(num))
	for i, ch := range num {
		val := int(ch - '0')
		if change[val] <= val {
			res = append(res, byte(ch))
			continue
		}
		// Replace for as long as it is an improvement
		j := i
		for ; j < len(num); j++ {
			val := int(num[j] - '0')
			if change[val] < val {
				break
			}
			res = append(res, byte(change[val]+'0'))
		}
		res = append(res, []byte(num)[j:]...)
		break
	}
	return string(res)
}
