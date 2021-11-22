package p2081sumofkmirrornumbers

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kMirror(t *testing.T) {
	for _, tc := range []struct {
		k    int
		n    int
		want int64
	}{
		{5, 25, 6849225412},
		{2, 5, 25},
		{3, 7, 499},
		{7, 17, 20379000},
	} {
		t.Run(fmt.Sprintf("%+v", tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, kMirror(tc.k, tc.n))
		})
	}
}

func kMirror(k int, n int) int64 {
	candidates := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		prefix := make([]int, 20)
		for width := 1; width < 20; width++ {
			if dfs(ctx, prefix, candidates, 0, width) {
				return
			}
		}
	}()

	var sum int
	baseVals := make([]int, 64)
	for n > 0 {
		cand := <-candidates
		if isMirrored(baseVals, cand, k) {
			n--
			sum += cand
		}
	}
	cancel()
	return int64(sum)
}

func isMirrored(baseVals []int, val, base int) bool {
	var pos int
	for val > 0 {
		baseVals[pos] = val % base
		pos++
		val /= base
	}
	for l, r := 0, pos-1; l < r; l, r = l+1, r-1 {
		if baseVals[l] != baseVals[r] {
			return false
		}
	}
	return true
}

func dfs(ctx context.Context, prefix []int, candidates chan int, pos int, width int) bool {
	if nnum := (width/2 + width%2); pos == nnum {
		var val int
		for i := 0; i < nnum; i++ {
			val *= 10
			val += prefix[i]
		}
		if width%2 == 0 {
			val *= 10
			val += prefix[nnum-1]
		}
		for i := nnum - 2; i >= 0; i-- {
			val *= 10
			val += prefix[i]
		}
		select {
		case <-ctx.Done():
			return true
		case candidates <- val:
			return false
		}
	}
	var start int
	if pos == 0 {
		start = 1
	}
	for i := start; i <= 9; i++ {
		prefix[pos] = i
		if dfs(ctx, prefix, candidates, pos+1, width) {
			return true
		}
	}
	return false
}
