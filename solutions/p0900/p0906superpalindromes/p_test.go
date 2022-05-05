package p0906superpalindromes

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_superpalindromesInRange(t *testing.T) {
	for _, tc := range []struct {
		left  string
		right string
		want  int
	}{
		{"398904669", "13479046850", 6},
		{"40000000000000000", "50000000000000000", 2},
		{"4", "1000", 4},
		{"1", "2", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.left), func(t *testing.T) {
			require.Equal(t, tc.want, superpalindromesInRange(tc.left, tc.right))
		})
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func superpalindromesInRange(left string, right string) int {
	lower, err := strconv.Atoi(left)
	check(err)
	upper, err := strconv.Atoi(right)
	check(err)

	// Odd palindrome
	superPalindromes := make(map[string]struct{})
	for i := 1; i < 100000; i++ {
		// Form square root palindromes by copying i to the left side
		s := strconv.Itoa(i)

		ss := s + rev(s[:len(s)-1])
		ssn, err := strconv.Atoi(ss)
		check(err)
		ssn *= ssn
		if ssn > upper {
			break
		}
		if ssn >= lower {
			sss := strconv.Itoa(ssn)
			if sss == rev(sss) {
				superPalindromes[sss] = struct{}{}
			}
		}
	}

	// Even palindrome
	for i := 1; i < 100000; i++ {
		// Form square root palindromes by copying i to the left side
		s := strconv.Itoa(i)

		ss := s + rev(s)

		// Check if square is also palindrome
		ssn, err := strconv.Atoi(ss)
		check(err)
		ssn *= ssn
		if ssn > upper {
			break
		}
		if ssn >= lower {
			sss := strconv.Itoa(ssn)
			if sss == rev(sss) {
				superPalindromes[sss] = struct{}{}
			}
		}
	}

	return len(superPalindromes)
}

// func isPalindrome(n int) bool {
// 	bs := make([]byte, 0)
// 	for n >= 10 {
// 		bs = append(bs, byte(n%10))
// 		n /= 10
// 	}
// 	bs = append(bs, byte(n))
// 	for l, r := 0, len(bs)-1; l <= r; l, r = l+1, r-1 {
// 		if bs[l] != bs[r] {
// 			return false
// 		}
// 	}
// 	return true
// }

func rev(s string) string {
	bs := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	return string(bs)
}
