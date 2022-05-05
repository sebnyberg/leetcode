package p0869reorderedpowerof2

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reorderedPowerOf2(t *testing.T) {
	for _, tc := range []struct {
		N    int
		want bool
	}{
		{1, true},
		{10, false},
		{16, true},
		{24, false},
		{46, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.N), func(t *testing.T) {
			require.Equal(t, tc.want, reorderedPowerOf2(tc.N))
		})
	}
}

// Version 1
// Calculate all powers of 2 (up to 32 bits)
// Convert to a sorted string of numbers
// Convert N to a sorted string and check if it is in the list of numbers
func reorderedPowerOf2(N int) bool {
	powersOf2 := make(map[string]struct{})
	n := 1
	for i := 0; i < 32; i++ {
		s := []byte(strconv.Itoa(n))
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		powersOf2[string(s)] = struct{}{}
		n <<= 1
	}

	s := []byte(strconv.Itoa(N))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	_, exists := powersOf2[string(s)]
	return exists
}
