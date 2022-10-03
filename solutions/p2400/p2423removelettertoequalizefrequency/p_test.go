package p2423removelettertoequalizefrequency

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_equalFrequency(t *testing.T) {
	for _, tc := range []struct {
		word string
		want bool
	}{
		{"abcc", true},
		{"aazz", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, equalFrequency(tc.word))
		})
	}
}

func equalFrequency(word string) bool {
	try := func(deletePos int) bool {
		var freq [26]int
		for i, ch := range word {
			if i == deletePos {
				continue
			}
			freq[ch-'a']++
		}
		var curr int
		for _, count := range freq {
			if count == 0 {
				continue
			}
			if curr == 0 {
				curr = count
				continue
			}
			if count != curr {
				return false
			}
		}
		return true
	}
	for i := range word {
		if try(i) {
			return true
		}
	}
	return false
}
