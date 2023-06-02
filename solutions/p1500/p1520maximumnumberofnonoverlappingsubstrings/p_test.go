package p1520maximumnumberofnonoverlappingsubstrings

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxNumOfSubstrings(t *testing.T) {
	for i, tc := range []struct {
		s    string
		want []string
	}{
		{"abab", []string{"abab"}},
		{"adefaddaccc", []string{"e", "f", "ccc"}},
		{"abbaccd", []string{"d", "bb", "cc"}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			res := maxNumOfSubstrings(tc.s)
			sort.Strings(tc.want)
			sort.Strings(res)
			require.Equal(t, tc.want, res)
		})
	}
}

func maxNumOfSubstrings(s string) []string {
	var left [26]int
	var right [26]int
	for i := range left {
		left[i] = math.MaxInt32
		right[i] = -1
	}
	for i, ch := range s {
		c := ch - 'a'
		left[c] = min(left[c], i)
		right[c] = max(right[c], i)
	}
	seqs := make([]sequence, 26)
	for i := range seqs {
		seqs[i].l = -1
	}
	for ch, i := range left {
		if i == math.MaxInt32 {
			continue
		}
		j := right[ch]
		seqs[ch] = sequence{
			ch: ch,
			l:  i,
			r:  j + 1,
		}
	}
	var count [26]int
	for i, ch := range s {
		ch -= 'a'
		if seqs[ch].l == i {
			seqs[ch].startCount = count
		}
		count[ch]++
		if seqs[ch].r == i+1 {
			seqs[ch].endCount = count
		}
	}
	sort.Slice(seqs, func(i, j int) bool {
		a := seqs[i]
		b := seqs[j]
		if a.l == -1 {
			return false
		}
		if b.l == -1 {
			return true
		}
		return a.l < b.l
	})
	// Prune -1s
	for i := range seqs {
		if seqs[i].l == -1 {
			seqs = seqs[:i]
			break
		}
	}

	// Find solution using DFS.
	ss := make([]string, 0, len(seqs))
	var d dfser
	d.dfs(s, ss, seqs, 0)
	return d.res
}

type dfser struct {
	res []string
	m   int
}

type sequence struct {
	ch         int
	l          int
	r          int
	startCount [26]int
	endCount   [26]int
}

func (d *dfser) dfs(s string, ss []string, seqs []sequence, i int) {
	if i == len(seqs) {
		var n int
		for _, s := range ss {
			n += len(s)
		}
		if len(ss) > len(d.res) || (len(ss) == len(d.res) && n < d.m) {
			d.m = n
			d.res = append(d.res[:0], ss...)
		}
		return
	}
	// Don't pick
	d.dfs(s, ss, seqs, i+1)

	// Pick
	//
	// The hypothesis is that the current letter is the start of a valid
	// sequence. For any new sequence started before reaching the end, we move
	// the end forward. If the final end count - start count includes a letter
	// that has not been seen, then there must exist letters inside the range
	// that are part of the end of a subarray that started before this one -
	// hence no valid sequence starts here.
	//
	j := i + 1
	r := seqs[i].r
	endCount := seqs[i].endCount
	seen := (1 << seqs[i].ch)
	for j < len(seqs) && seqs[j].l < r {
		if seqs[j].r > r {
			r = seqs[j].r
			endCount = seqs[j].endCount
		}
		seen |= (1 << seqs[j].ch)
		j++
	}
	// Verify that no letter exists in the range that wasn't started as well.
	for ch := range seqs[i].startCount {
		d := endCount[ch] - seqs[i].startCount[ch]
		if d > 0 && seen&(1<<ch) == 0 {
			// Not valid
			return
		}
	}
	t := s[seqs[i].l:r]
	ss = append(ss, t)
	d.dfs(s, ss, seqs, j)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
