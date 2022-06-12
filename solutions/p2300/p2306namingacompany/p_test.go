package p2306namingacompany

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distinctNames(t *testing.T) {
	for _, tc := range []struct {
		ideas []string
		want  int64
	}{
		{[]string{"coffee", "donuts", "time", "toffee"}, 6},
		{[]string{"lack", "back"}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.ideas), func(t *testing.T) {
			require.Equal(t, tc.want, distinctNames(tc.ideas))
		})
	}
}

func distinctNames(ideas []string) int64 {
	strIdx := make(map[string]int)
	letters := make([]uint32, 0, 1000)
	for _, idea := range ideas {
		k := idea[1:]
		if _, exists := strIdx[k]; !exists {
			strIdx[k] = len(letters)
			letters = append(letters, 0)
		}
		letters[strIdx[k]] |= 1 << (idea[0] - 'a')
	}

	// For each combination of words, count overlapping first letters
	var res int
	for i, l1 := range letters {
		for j := i + 1; j < len(letters); j++ {
			l2 := letters[j]
			m := ^(l1 & l2)
			res += bits.OnesCount32(l1&m) * bits.OnesCount32(l2&m)
		}
	}

	return int64(res * 2)
}

// func distinctNames(ideas []string) int64 {
// 	// Group by letter
// 	var letterStrs [26]map[string]struct{}
// 	for i := 0; i < 26; i++ {
// 		letterStrs[i] = make(map[string]struct{}, 1000)
// 	}
// 	for _, idea := range ideas {
// 		letterStrs[idea[0]-'a'][idea[1:]] = struct{}{}
// 	}

// 	// For each combination of letters
// 	var res int
// 	for a := 0; a < 26-1; a++ {
// 		for b := a + 1; b < 26; b++ {
// 			var count [2]int
// 			// For each postfix that starts with a
// 			for postfix := range letterStrs[a] {
// 				// If the postfix does not exist for b
// 				if _, exists := letterStrs[b][postfix]; !exists {
// 					// Then b + postfix is valid
// 					count[0]++
// 				}
// 			}
// 			// Vice versa for a+postfix
// 			for postfix := range letterStrs[b] {
// 				if _, exists := letterStrs[a][postfix]; !exists {
// 					count[1]++
// 				}
// 			}
// 			res += count[0] * count[1]
// 		}
// 	}

// 	return int64(res * 2)
// }
