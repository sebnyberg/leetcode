package p1419minimumnumberoffrogscroaking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minNumberOfFrogs(t *testing.T) {
	for i, tc := range []struct {
		croakOfFrogs string
		want         int
	}{
		{"croakcroak", 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minNumberOfFrogs(tc.croakOfFrogs))
		})
	}
}

func minNumberOfFrogs(croakOfFrogs string) int {
	// Because each letter is unique, we can count occurrences and move each
	// index forward.
	count := make(map[rune]int)
	prev := []rune{
		'r': 'c',
		'o': 'r',
		'a': 'o',
		'k': 'a',
	}
	var nparallel int
	var res int
	for _, ch := range croakOfFrogs {
		if ch == 'c' {
			nparallel++
			res = max(res, nparallel)
			count['c']++
			continue
		}
		pre := prev[ch]
		if count[pre] == 0 {
			return -1
		}
		count[ch]++
		count[pre]--
		if ch == 'k' {
			nparallel--
		}
	}
	for _, ch := range "croa" {
		if count[ch] > 0 {
			return -1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
