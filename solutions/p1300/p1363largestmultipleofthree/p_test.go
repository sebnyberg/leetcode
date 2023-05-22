package p1363largestmultipleofthree

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestMultipleOfThree(t *testing.T) {
	for i, tc := range []struct {
		digits []int
		want   string
	}{
		{[]int{9, 7, 6, 7, 6}, ""},
		{[]int{5, 8}, ""},
		{[]int{8, 1, 9}, "981"},
		{[]int{8, 6, 7, 1, 0}, "8760"},
		{[]int{1}, ""},
		{[]int{0, 0, 0, 0}, "0"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, largestMultipleOfThree(tc.digits))
		})
	}
}

func largestMultipleOfThree(digits []int) string {
	// Weird fact about multiples of three: if the sum of the digits is
	// divisible by three, then so is the number. I recall a teacher blowing my
	// mind with this in high-school :)
	//
	// If we add all digits together, we get a certain sum % 3
	var minMod21, minMod22 int = math.MaxInt32, math.MaxInt32
	var minMod11, minMod12 int = math.MaxInt32, math.MaxInt32
	var rest int
	var count [10]int
	for _, x := range digits {
		rest = (rest + x) % 3
		if x%3 == 1 {
			if x <= minMod11 {
				minMod11, minMod12 = x, minMod11
			} else if x <= minMod12 {
				minMod12 = x
			}
		} else if x%3 == 2 {
			if x <= minMod21 {
				minMod21, minMod22 = x, minMod21
			} else if x <= minMod22 {
				minMod22 = x
			}
		}
		count[x]++
	}
	if rest == 1 {
		// Remove smallest %3==1 number, or two smallest %3==2
		if minMod11 != math.MaxInt32 {
			count[minMod11]--
		} else {
			count[minMod21]--
			count[minMod22]--
		}
	} else if rest == 2 {
		// Remove smallest %3==2 number, or two smallest %3==1
		if minMod21 != math.MaxInt32 {
			count[minMod21]--
		} else {
			count[minMod11]--
			count[minMod12]--
		}
	}
	var bs []byte
	for x := 9; x >= 0; x-- {
		c := count[x]
		for i := 0; i < c; i++ {
			bs = append(bs, byte('0'+x))
		}
	}
	res := string(bs)
	if len(res) > 1 {
		res = strings.TrimLeft(res[:len(res)-1], "0") + string(res[len(res)-1])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
