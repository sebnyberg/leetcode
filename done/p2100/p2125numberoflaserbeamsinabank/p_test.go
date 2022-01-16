package p2025numberoflaserbeamsinabank

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfBeams(t *testing.T) {
	for _, tc := range []struct {
		bank []string
		want int
	}{
		{[]string{"011001", "000000", "010100", "001000"}, 8},
		{[]string{"000", "111", "000"}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.bank), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfBeams(tc.bank))
		})
	}
}

func numberOfBeams(bank []string) int {
	var prevCount int
	var res int
	for _, row := range bank {
		if strings.ContainsRune(row, '1') {
			cnt := strings.Count(row, "1")
			res += prevCount * cnt
			prevCount = cnt
		}
	}
	return res
}
