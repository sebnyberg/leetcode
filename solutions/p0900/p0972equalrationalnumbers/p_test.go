package p0972equalrationalnumbers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isRationalEqual(t *testing.T) {
	for i, tc := range []struct {
		s    string
		t    string
		want bool
	}{
		{"1.0", "1.", true},
		{"0.9(9)", "1.", true},
		{"0.(52)", "0.5(25)", true},
		{"1.9(0)", "1.8(9)", true},
		{"0.1666(6)", "0.166(66)", true},
		{"8.123(4567)", "8.123(4566)", false},
		{"1.001(01)", "1.00(10)", true},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, isRationalEqual(tc.s, tc.t))
		})
	}
}

var pattern = regexp.MustCompile(`([0-9]+)\.?([0-9]+)?(\(([0-9]+)\))?`)

func isRationalEqual(s string, t string) bool {
	normalize := func(s string) string {
		// Scan input
		parts := pattern.FindStringSubmatch(s)
		a, _ := strconv.Atoi(parts[1])

		dec := parts[2]
		repeat := parts[4]

		// If there's a repeating part, repeat a bunch of times.
		if strings.Trim(repeat, "0") != "" {
			for len(dec) < 20 {
				dec += repeat
			}
			if len(dec) > 20 {
				dec = dec[:20]
			}

			// Round up repeating 9s
			if strings.TrimRight(repeat, "9") == "" {
				dec = strings.TrimRight(dec, "9")

				// Either increment the natural number, or add the carry to the
				// last number of the decimal
				m := len(dec)
				if m == 0 {
					return fmt.Sprintf("%d.", a+1)
				}
				return fmt.Sprintf("%d.%s%c", a, dec[:m-1], dec[m-1]+1)
			}
		}

		// Trim trailing zeroes
		return fmt.Sprintf("%d.%s", a, strings.TrimRight(dec, "0"))
	}
	s = normalize(s)
	t = normalize(t)
	return s == t
}
