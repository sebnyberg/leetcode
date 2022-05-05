package p1805numberofdifferentintsinstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numDifferentIntegers(t *testing.T) {
	for _, tc := range []struct {
		word string
		want int
	}{
		{"4w31am0ets6sl5go5ufytjtjpb7b0sxqbee2blg9ss", 8},
		{"a1b01c001", 1},
		{"a123bc34d8ef34", 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, numDifferentIntegers(tc.word))
		})
	}
}

func numDifferentIntegers(word string) int {
	nums := make(map[string]struct{})
	res := 0
	n := len(word)
	for i := 0; i < n; {
		d := word[i] - '0'
		if d >= 10 {
			i++
			continue
		}
		l := i
		i++
		for i < n && word[i]-'0' < 10 {
			i++
		}
		for l < n && word[l] == '0' && l < i {
			l++
		}
		k := word[l:i]
		if _, exists := nums[k]; !exists {
			res++
			nums[k] = struct{}{}
		}
	}
	return res
}
