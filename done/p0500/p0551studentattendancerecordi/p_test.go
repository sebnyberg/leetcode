package p0551studentattendancerecordi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkRecord(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"LLL", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, checkRecord(tc.s))
		})
	}
}

func checkRecord(s string) bool {
	var lateCount int
	var absenceCount int
	for i := range s {
		if s[i] == 'A' {
			absenceCount++
		}
		if absenceCount == 2 {
			return false
		}
		if s[i] == 'L' {
			if i > 0 && s[i-1] == 'L' {
				lateCount++
			} else {
				lateCount = 1
			}
			if lateCount == 3 {
				return false
			}
		}
	}
	return true
}
