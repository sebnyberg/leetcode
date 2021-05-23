package p1871jumpgamevii

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canReach(t *testing.T) {
	for _, tc := range []struct {
		s       string
		minJump int
		maxJump int
		want    bool
	}{
		{"011010", 2, 3, true},
		{"01101110", 2, 3, false},
		{"00111010", 3, 5, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, canReach(tc.s, tc.minJump, tc.maxJump))
		})
	}
	f, _ := os.Open("input")
	sc := bufio.NewScanner(f)
	sc.Scan()
	ss, _ := io.ReadAll(f)
	sss := strings.Split(string(ss), "\n")
	s := sss[0]
	minJump, _ := strconv.Atoi(sss[1])
	maxJump, _ := strconv.Atoi(sss[2])
	res := canReach(s, minJump, maxJump)
	require.Equal(t, true, res)
}

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	ok := make([]bool, n)
	for i, ch := range s {
		ok[i] = ch == '0'
	}
	return canReachHelper(ok, minJump, maxJump, n, 0)
}

func canReachHelper(ok []bool, minJump, maxJump, n, pos int) bool {
	if pos+minJump <= n-1 && pos+maxJump >= n-1 {
		return true
	}
	for i := minJump; i <= min(n-1-pos, maxJump); i++ {
		if !ok[pos+i] {
			continue
		}
		if canReachHelper(ok, minJump, maxJump, n, pos+i) {
			return true
		}
	}
	ok[pos] = false
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
