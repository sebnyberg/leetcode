package p0906superpalindromes

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_superpalindromesInRange(t *testing.T) {
	for _, tc := range []struct {
		left  string
		right string
		want  int
	}{
		{"398904669", "13479046850", 6},
		{"40000000000000000", "50000000000000000", 2},
		{"4", "1000", 4},
		{"1", "2", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.left), func(t *testing.T) {
			require.Equal(t, tc.want, superpalindromesInRange(tc.left, tc.right))
		})
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func superpalindromesInRange(left string, right string) int {
	l := mustParse(left)
	r := mustParse(right)
	rev := func(s string) string {
		b := []byte(s)
		for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
			b[l], b[r] = b[r], b[l]
		}
		return string(b)
	}
	var res int
	for k := 1; ; k++ {
		if k == 1 {
			for x := 1; x <= 9; x++ {
				if x*x < l {
					continue
				}
				if x*x > r {
					return res
				}
				if isPalin(x * x) {
					res++
				}
			}
			continue
		}
		lo := pow(10, k/2-1)
		hi := pow(10, k/2)
		for y := 0; y < pow(10, k&1); y++ {
			for x := lo; x < hi; x++ {
				var mid string
				if k&1 == 1 {
					mid = fmt.Sprint(y)
				}
				v := mustParse(fmt.Sprint(x) + mid + rev(fmt.Sprint(x)))
				if v*v < l {
					continue
				}
				if v*v > r {
					return res
				}
				if isPalin(v * v) {
					res++
				}
			}
		}
	}
}

func mustParse(s string) int {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(v)
}

func isPalin(v int) bool {
	s := fmt.Sprint(v)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	return pow(a, b/2) * pow(a, b/2+b&1)
}
