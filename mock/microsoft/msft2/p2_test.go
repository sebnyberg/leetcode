package msft2_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_compareVersion(t *testing.T) {
	for _, tc := range []struct {
		version1 string
		version2 string
		want     int
	}{
		{"1.01", "1.001", 0},
		{"1.0", "1.0.0", 0},
		{"0.1", "1.0", -1},
		{"1.0.1", "1", 1},
		{"7.5.2.4", "7.5.3", -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.version1), func(t *testing.T) {
			require.Equal(t, tc.want, compareVersion(tc.version1, tc.version2))
		})
	}
}

func compareVersion(version1 string, version2 string) int {
	parts1 := strings.Split(version1, ".")
	parts2 := strings.Split(version2, ".")
	switched := false
	if len(parts1) < len(parts2) {
		parts1, parts2 = parts2, parts1
		switched = true
	}
	n2 := len(parts2)

	var res int
	for i, v1str := range parts1 {
		var v2str string
		if i >= n2 {
			v2str = "0"
		} else {
			v2str = parts2[i]
		}
		v1, _ := strconv.Atoi(v1str)
		v2, _ := strconv.Atoi(v2str)
		if v1 == v2 {
			continue
		}
		if v1 > v2 {
			res = 1
		} else {
			res = -1
		}
		break
	}

	// Switch back
	if switched {
		if res == -1 {
			return 1
		} else if res == 1 {
			return -1
		}
	}
	return res
}
