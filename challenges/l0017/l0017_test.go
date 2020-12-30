package l0017_test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func Test_letterCombinations(t *testing.T) {
// 	for _, tc := range []struct {
// 		in   string
// 		want []string
// 	}{
// 		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
// 		{"", []string{}},
// 		{"2", []string{"a", "b", "c"}},
// 	} {
// 		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
// 			require.Equal(t, tc.want, letterCombinations(tc.in))
// 		})
// 	}
// }

// const (
// 	two   = "abc"
// 	three = "def"
// 	four  = "ghi"
// 	five  = "jkl"
// 	six   = "mno"
// 	seven = "pqrs"
// 	eight = "tuv"
// 	nine  = "wxyz"
// )

// // Return all possible letter combinations that the number could represent
// // in any order.
// func letterCombinations(digits string) []string {
// 	if len(digits) == 0 {
// 		return []string{}
// 	}

// 	// Calculate size of result
// 	size := 1
// 	for _, digit := range digits {
// 		switch digit {
// 		case '2', '3', '4', '5', '6', '8':
// 			size *= 3
// 		default:
// 			size *= 4
// 		}
// 	}
// 	res := make([]string, size)
// 	for _, digit := range digits {

// 	}

// 	return res
// }
