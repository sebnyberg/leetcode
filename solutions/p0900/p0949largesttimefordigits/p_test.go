package p0949largesttimefordigits

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestTimeFromDigits(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want string
	}{
		{[]int{2, 4, 7, 9}, ""},
		{[]int{5, 5, 5, 5}, ""},
		{[]int{0, 4, 0, 0}, "04:00"},
		{[]int{1, 2, 3, 4}, "23:41"},
		{[]int{0, 0, 0, 0}, "00:00"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, largestTimeFromDigits(tc.arr))
		})
	}
}

func largestTimeFromDigits(arr []int) string {
	nums := make(map[int]int)
	for _, n := range arr {
		nums[n]++
	}

	findBelow := func(below int) int {
		val := -1
		for n, c := range nums {
			if c == 0 {
				continue
			}
			val = maxBelow(val, n, below)
		}
		return val
	}

	first := findBelow(3)
	switch {
	case first == -1:
		return ""
	case first == 2:
		// check if remainder is valid with 2 in the first slot
		nums[first]--
		underFour := findBelow(4)
		nums[underFour]--
		underSix := findBelow(6)
		nums[underSix]--
		underTen := findBelow(10)
		nums[underTen]--
		if underFour == -1 || underSix == -1 || underTen == -1 { // invalid with 2 in the first slot, continue
			nums[first]++
			nums[underFour]++
			nums[underSix]++
			nums[underTen]++
			// first must be below 2
			first = findBelow(2)
			if first == -1 {
				return ""
			}
			break
		}

		// Number was valid! Write to res and quit
		return fmt.Sprintf("%v%v:%v%v", first, underFour, underSix, underTen)
	}
	nums[first]--

	// Write first number
	res := strconv.Itoa(first)

	// Find max below 10 (cannot fail)
	underTen := findBelow(10)
	nums[underTen]--
	res += strconv.Itoa(underTen) + ":"

	// Max below 6
	underSix := findBelow(6)
	if underSix == -1 {
		return ""
	}
	nums[underSix]--
	res += strconv.Itoa(underSix)

	// Last cannot fail
	res += strconv.Itoa(findBelow(10))

	return res
}

func maxBelow(a, b, below int) int {
	if b >= below {
		return a
	}
	if b > a {
		return b
	}
	return a
}
