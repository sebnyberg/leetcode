package p0843guesstheword

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_a(t *testing.T) {
	for _, tc := range []struct {
		word    string
		guesses int
		words   []string
	}{
		{"hbaczn", 10, []string{"gaxckt", "trlccr", "jxwhkz", "ycbfps", "peayuf", "yiejjw", "ldzccp", "nqsjoa", "qrjasy", "pcldos", "acrtag", "buyeia", "ubmtpj", "drtclz", "zqderp", "snywek", "caoztp", "ibpghw", "evtkhl", "bhpfla", "ymqhxk", "qkvipb", "tvmued", "rvbass", "axeasm", "qolsjg", "roswcb", "vdjgxx", "bugbyv", "zipjpc", "tamszl", "osdifo", "dvxlxm", "iwmyfb", "wmnwhe", "hslnop", "nkrfwn", "puvgve", "rqsqpq", "jwoswl", "tittgf", "evqsqe", "aishiv", "pmwovj", "sorbte", "hbaczn", "coifed", "hrctvp", "vkytbw", "dizcxz", "arabol", "uywurk", "ppywdo", "resfls", "tmoliy", "etriev", "oanvlx", "wcsnzy", "loufkw", "onnwcy", "novblw", "mtxgwe", "rgrdbt", "ckolob", "kxnflb", "phonmg", "egcdab", "cykndr", "lkzobv", "ifwmwp", "jqmbib", "mypnvf", "lnrgnj", "clijwa", "kiioqr", "syzebr", "rqsmhg", "sczjmz", "hsdjfp", "mjcgvm", "ajotcx", "olgnfv", "mjyjxj", "wzgbmg", "lpcnbj", "yjjlwn", "blrogv", "bdplzs", "oxblph", "twejel", "rupapy", "euwrrz", "apiqzu", "ydcroj", "ldvzgq", "zailgu", "xgqpsr", "wxdyho", "alrplq", "brklfk"}},
		{"vftnkr", 12, []string{"mjpsce", "giwiyk", "slbnia", "pullbr", "ezvczd", "dwkrmt", "qgzebh", "wvhhlm", "kqbmny", "zpvrkz", "pdwxvy", "gilywa", "gmrrdc", "vvqvla", "rmjirt", "qmvykq", "mhbmuq", "unplzn", "qkcied", "eignxg", "fbfgng", "xpizga", "twubzr", "nnfaxr", "skknhe", "twautl", "nglrst", "mibyks", "qrbmpx", "ukgjkq", "mhxxfb", "deggal", "bwpvsp", "uirtak", "tqkzfk", "hfzawa", "jahjgn", "mteyut", "jzbqbv", "ttddtf", "auuwgn", "untihn", "gbhnog", "zowaol", "feitjl", "omtiur", "kwdsgx", "tggcqq", "qachdn", "dixtat", "hcsvbw", "chduyy", "gpdtft", "bjxzky", "uvvvko", "jzcpiv", "gtyjau", "unsmok", "vfcmhc", "hvxnut", "orlwku", "ejllrv", "jbrskt", "xnxxdi", "rfreiv", "njbvwj", "pkydxy", "jksiwj", "iaembk", "pyqdip", "exkykx", "uxgecc", "khzqgy", "dehkbu", "ahplng", "jomiik", "nmcsfe", "bclcbp", "xfiefi", "soiwde", "tcjkjp", "wervlz", "dcthgv", "hwwghe", "hdlkll", "dpzoxb", "mpiviy", "cprcwo", "molttv", "dwjtdp", "qiilsr", "dbnaxs", "dbozaw", "webcyp", "vftnkr", "iurlzf", "giqcfc", "pcghoi", "zupyzn", "xckegy"}},
		{"azzzzz", 10, []string{"abcdef", "acdefg", "adefgh", "aefghi", "afghij", "aghijk", "ahijkl", "aijklm", "ajklmn", "aklmno", "almnoz", "anopqr", "azzzzz"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			m := &Master{correctWord: tc.word, remainingGuesses: tc.guesses}
			findSecretWord(tc.words, m)
			require.True(t, m.didGuessRight)
		})
	}
}

type Master struct {
	correctWord      string
	didGuessRight    bool
	remainingGuesses int
}

func (m *Master) Guess(word string) int {
	if m.didGuessRight {
		return -1
	}
	if m.remainingGuesses == 0 {
		m.didGuessRight = false
		return -1
	}
	var res int
	for i := range word {
		if word[i] == m.correctWord[i] {
			res++
		}
	}
	if word == m.correctWord {
		m.didGuessRight = true
	}
	m.remainingGuesses--
	return res
}

func findSecretWord(words []string, master *Master) {
	// We can guess at least 10 times, so there's that.
	// Before that, we want to eliminate as many options as possible.
	// Let's compare two words, "aabbcc" and "ababdd"
	// The two words share 2 character/position pairs, i.e. "a" in position 0 and
	// "b" in position 3.
	// So if we pass in "aabbcc", and the matcher returns anything aside from 2,
	// then we can rule out the second word.
	//
	// This means that we can maintain a list of number of matching elements for
	// each word that we guess. The best guess is a word which has a pretty even
	// distribution of shared character/position pairs. For example, a word that
	// shares 1 letter with 1/5 of the words, 2 letters with 1/5, and so on, would
	// be great because it rules out 4/5ths of all words.
	n := len(words)
	invalid := make([]bool, n)
	nleft := n
	inCommon := func(a, b string) int {
		var res int
		for i := range a {
			if a[i] == b[i] {
				res++
			}
		}
		return res
	}
	calcScore := func(mappings [6][]int) int {
		// Median score of a non-zero length
		as := []int{}
		for _, m := range mappings {
			n := len(m)
			if n == 0 {
				continue
			}
			as = append(as, n)
		}
		if len(as) == 1 {
			return 0
		}
		mid := len(as) / 2
		if len(as)&1 == 1 {
			return as[mid]
		}
		return min(as[mid-1], as[mid])
	}
	guessesLeft := 10
	for nleft > guessesLeft {
		var maxScoreWord int
		var maxScoreMappings [6][]int
		maxScore := -1
		for i := 0; i < n; i++ {
			var inCommonMapping [6][]int
			if invalid[i] {
				continue
			}
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				if invalid[j] {
					continue
				}
				x := inCommon(words[i], words[j])
				inCommonMapping[x] = append(inCommonMapping[x], j)
			}
			score := calcScore(inCommonMapping)
			if score > maxScore {
				maxScoreWord = i
				maxScore = score
				for x := 0; x < 6; x++ {
					maxScoreMappings[x] = maxScoreMappings[x][:0]
					maxScoreMappings[x] = append(maxScoreMappings[x], inCommonMapping[x]...)
				}
			}
		}
		// Guess best word
		w := words[maxScoreWord]
		res := master.Guess(w)
		if res == 6 {
			return
		}
		invalid[maxScoreWord] = true
		nleft--
		// Remove all words that don't match the result
		for x := 0; x < 6; x++ {
			if res == x {
				continue
			}
			for _, j := range maxScoreMappings[x] {
				invalid[j] = true
				nleft--
			}
		}
	}
	// Guess remaining words
	for i := range words {
		if invalid[i] {
			continue
		}
		if res := master.Guess(words[i]); res == 6 {
			return
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
