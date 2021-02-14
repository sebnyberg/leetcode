package p0093restoreipaddr

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_restoreIpAddresses(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want []string
	}{
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"0000", []string{"0.0.0.0"}},
		{"1111", []string{"1.1.1.1"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, restoreIpAddresses(tc.s))
		})
	}
}

var maxSize = 1<<8 - 1

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	findIPAddr(s, []byte{}, &res)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findIPAddr(s string, prefix []byte, res *[]string) {
	// Add valid IP addrs
	if len(prefix) > 4 {
		return
	}
	if len(s) == 0 {
		if len(prefix) == 4 {
			ipStr := make([]string, 4)
			for i, b := range prefix {
				ipStr[i] = strconv.Itoa(int(b))
			}
			*res = append(*res, strings.Join(ipStr, "."))
		}
		return
	}

	switch s[0] {
	case '0': // only one option, take the zero
		prefix = append(prefix, 0)
		findIPAddr(s[1:], prefix, res)
		return
	default:
		for i := 0; i < min(len(s), 3); i++ {
			n, _ := strconv.Atoi(s[:i+1])
			if n > maxSize {
				break
			}
			prefixCpy := make([]byte, len(prefix))
			copy(prefixCpy, prefix)
			prefixCpy = append(prefixCpy, byte(n))
			findIPAddr(s[i+1:], prefixCpy, res)
		}
	}
}
