package p0068textjustification

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fullJustify(t *testing.T) {
	for _, tc := range []struct {
		words    []string
		maxWidth int
		want     []string
	}{
		{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			20, []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			}},
		{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16, []string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		}},
		{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16, []string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		}},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.words, tc.maxWidth), func(t *testing.T) {
			require.Equal(t, tc.want, fullJustify(tc.words, tc.maxWidth))
		})
	}
}

func fullJustify(words []string, maxWidth int) []string {
	lineWords := make([][]string, 1)
	var i int
	curLen := len(words[0])
	lineWords[i] = []string{words[0]}
	for _, word := range words[1:] {
		wordLen := len(word)
		if curLen+wordLen+1 > maxWidth {
			// Roll over to next line
			lineWords = append(lineWords, []string{word})
			curLen = wordLen
			i++
			continue
		}
		lineWords[i] = append(lineWords[i], word)
		curLen += wordLen + 1
	}

	return justifyLines(lineWords, maxWidth)
}

func justifyLines(lineWords [][]string, maxWidth int) []string {
	n := len(lineWords)
	justifiedLines := make([]string, n)
	for i := 0; i < n-1; i++ {
		var lineLen int
		for _, word := range lineWords[i] {
			lineLen += len(word)
		}
		spaces := maxWidth - lineLen
		// If there is only one word, add spacing to the right and continue
		if len(lineWords[i]) == 1 {
			justifiedLines[i] = lineWords[i][0] + strings.Repeat(" ", spaces)
			continue
		}

		inbetweens := len(lineWords[i]) - 1
		inbetweenSpaces := spaces / inbetweens
		// If there is an uneven amount of inbetween spaces,
		// add an extra spaces to the end of each word
		// We know this is safe because the number of left-over spaces
		// is bounded by the number of words
		extraSpaces := spaces % inbetweens
		for j := 0; j < extraSpaces; j++ {
			lineWords[i][j] += " "
		}
		justifiedLines[i] = strings.Join(lineWords[i], strings.Repeat(" ", inbetweenSpaces))
	}

	// For the last line, join the words and add missing spaces
	justifiedLines[n-1] = strings.Join(lineWords[n-1], " ")
	justifiedLines[n-1] = justifiedLines[n-1] + strings.Repeat(" ", maxWidth-len(justifiedLines[n-1]))
	return justifiedLines
}
