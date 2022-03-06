package p0564findtheclosestpalindrome

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nearestPalindromic(t *testing.T) {
	for _, tc := range []struct {
		n    string
		want string
	}{
		{"99", "101"},
		{"100", "99"},
		{"11", "9"},
		{"10", "9"},
		{"1213", "1221"},
		{"123", "121"},
		{"1", "0"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, nearestPalindromic(tc.n))
		})
	}
}

func nearestPalindromic(n string) string {
	m := len(n)
	s := n
	num := func(s string) int {
		v, _ := strconv.Atoi(s)
		return v
	}
	mirror := s[:m/2] + s[m/2:m/2+m&1] + rev(s[:m/2])
	candidates := []int{num(mirror)}
	midL, midR := m/2-1+m&1, m/2
	if mirror[midL] != '0' {
		bs := []byte(mirror)
		bs[midL] = mirror[midL] - 1
		bs[midR] = mirror[midL] - 1
		candidates = append(candidates, num(string(bs)))
	}
	if mirror[midL] != '9' {
		bs := []byte(mirror)
		bs[midL] = mirror[midL] + 1
		bs[midR] = mirror[midL] + 1
		candidates = append(candidates, num(string(bs)))
	}
	candidates = append(candidates, pow(10, m-1)-1, pow(10, m-1)+1)
	candidates = append(candidates, pow(10, m)-1, pow(10, m)+1)
	orig := num(s)
	res := math.MaxInt32
	minDiff := math.MaxInt32
	for _, cand := range candidates {
		if cand == orig {
			continue
		}
		d := abs(cand - orig)
		if d < minDiff || d == minDiff && cand < res {
			res = cand
			minDiff = d
		}
	}
	return fmt.Sprint(res)
}

func pow(a, b int) int {
	res := a
	for i := 2; i <= b; i++ {
		res *= a
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func rev(s string) string {
	bs := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	return string(bs)
}
