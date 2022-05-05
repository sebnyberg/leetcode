package p0418sentencescreenfitting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordsTyping(t *testing.T) {
	for _, tc := range []struct {
		sentence   []string
		rows, cols int
		want       int
	}{
		{[]string{"i", "had", "apple", "pie"}, 4, 5, 1},
		{[]string{"a", "bcd", "e"}, 3, 6, 2},
		{[]string{"hello", "world"}, 2, 8, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.sentence), func(t *testing.T) {
			require.Equal(t, tc.want, wordsTyping(tc.sentence, tc.rows, tc.cols))
		})
	}
}

func wordsTyping(sentence []string, rows int, cols int) int {
	// dp[i] = number of words that fit starting with word with index i
	n := len(sentence)
	dp := make([]int, n)
	for i := 0; i < len(sentence); i++ {
		if cols < len(sentence[i]) {
			continue
		}
		remains := cols - len(sentence[i])
		count := 1
		for j := (i + 1) % n; remains >= len(sentence[j])+1; j = (j + 1) % n {
			remains -= len(sentence[j]) + 1
			count++
		}
		dp[i] = count
	}
	var dpIdx int
	var count int
	for row := 0; row < rows; row++ {
		nwords := dp[dpIdx%n]
		count += nwords
		dpIdx += nwords
	}
	return count / n
}
