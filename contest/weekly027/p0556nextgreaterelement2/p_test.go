package p0556nextgreaterelement2

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextGreaterElement(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{1234, 1243},
		{101, 110},
		{12, 21},
		{21, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, nextGreaterElement(tc.n))
		})
	}
}

func nextGreaterElement(n int) int {
	var numCount [10]int
	s := []byte(fmt.Sprint(n))
	// For each index i := 0; i < len(s)-1, find the smallest number prior to i which
	// is greater in size.
	numCount[s[len(s)-1]-'0']++
	for i := len(s) - 2; i >= 0; i-- {
		for j := s[i] - '0' + 1; j <= 9; j++ {
			if numCount[j] == 0 {
				continue
			}
			// Found a match. Put 'j' into the current position, then create the
			// smallest possible number after this position
			numCount[s[i]-'0']++
			numCount[j]--
			s[i] = j + '0'
			k := i + 1
			for m := 0; m <= 9; m++ {
				for n := 0; n < numCount[m]; n++ {
					s[k] = byte(m + '0')
					k++
				}
			}
			val, _ := strconv.Atoi(string(s))
			if val > math.MaxInt32 {
				return -1
			}
			return val
		}
		numCount[s[i]-'0']++
	}
	return -1
}
