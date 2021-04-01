package google1

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

func TestFenwick(t *testing.T) {
	na := Constructor([]int{1, 3, 5})
	require.Equal(t, 9, na.SumRange(0, 2))
	na.Update(1, 2)
	require.Equal(t, 8, na.SumRange(0, 2))
}

type NumArray struct {
	tree []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{
		tree: make([]int, len(nums)),
	}
	copy(na.tree, nums)
	for i := range na.tree {
		if j := i | (i + 1); j < len(nums) {
			na.tree[j] += na.tree[i]
		}
	}
	return na
}

func (this *NumArray) get(index int) int {
	sum := this.tree[index]
	j := index + 1
	j -= j & -j
	for index > j {
		sum -= this.tree[index-1]
		index -= index & -index
	}
	return sum
}

// Sum returns the sum of the elements from index 0 to index i-1.
func (this *NumArray) sum(index int) int {
	var sum int
	for index > 0 {
		sum += this.tree[index-1]
		index -= index & -index
	}
	return sum
}

func (this *NumArray) Update(index int, val int) {
	val -= this.get(index)
	for len := len(this.tree); index < len; index |= index + 1 {
		this.tree[index] += val
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	d := this.sum(right+1) - this.sum(left)
	return d
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

func Test_backspaceCompare(t *testing.T) {
	for _, tc := range []struct {
		S    string
		T    string
		want bool
	}{
		// {"ab#c", "ad#c", true},
		// {"ab##", "c#d#", true},
		// {"a##c", "#a#c", true},
		{"a#c", "b", false},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.S, tc.T), func(t *testing.T) {
			require.Equal(t, tc.want, backspaceCompare(tc.S, tc.T))
		})
	}
}

func backspaceCompare(S string, T string) bool {
	sstack := make([]byte, 0)
	for i := range S {
		if S[i] == '#' {
			if len(sstack) > 0 {
				sstack = sstack[:len(sstack)-1]
			}
		} else {
			sstack = append(sstack, S[i])
		}
	}
	tstack := make([]byte, 0)
	for i := range T {
		if T[i] == '#' {
			if len(tstack) > 0 {
				tstack = tstack[:len(tstack)-1]
			}
		} else {
			tstack = append(tstack, T[i])
		}
	}
	if len(sstack) != len(tstack) {
		return false
	}
	for i := range sstack {
		if sstack[i] != tstack[i] {
			return false
		}
	}
	return true
}

func Test_longestLine(t *testing.T) {
	for _, tc := range []struct {
		M    [][]int
		want int
	}{
		{[][]int{
			{0, 0, 1, 1},
			{0, 0, 1, 0},
			{0, 1, 0, 1},
		}, 3},
		{[][]int{
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 1},
		}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.M), func(t *testing.T) {
			require.Equal(t, tc.want, longestLine(tc.M))
		})
	}
}

func longestLine(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	m, n := len(M), len(M[0])
	var maxRunLen int

	// rows
	for i := range M {
		var curLen int
		for j := range M[i] {
			switch {
			case M[i][j] == 1:
				curLen++
				maxRunLen = max(maxRunLen, curLen)
			case M[i][j] == 0:
				curLen = 0
			}
		}
	}

	// cols
	for j := range M[0] {
		var curLen int
		for i := range M {
			switch {
			case M[i][j] == 1:
				curLen++
				maxRunLen = max(maxRunLen, curLen)
			case M[i][j] == 0:
				curLen = 0
			}
		}
	}

	validIndex := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	// Diagonal
	for j := -m + 1; j < n; j++ {
		var curLen int
		for k := 0; k < m; k++ {
			if !validIndex(k, j+k) {
				continue
			}
			switch {
			case M[k][j+k] == 1:
				curLen++
				maxRunLen = max(maxRunLen, curLen)
			case M[k][j+k] == 0:
				curLen = 0
			}
		}
	}

	// Antidiagonal
	for j := n - 1 + m - 1; j >= 0; j-- {
		var curLen int
		for k := 0; k < m; k++ {
			if !validIndex(k, j-k) {
				continue
			}
			switch {
			case M[k][j-k] == 1:
				curLen++
				maxRunLen = max(maxRunLen, curLen)
			case M[k][j-k] == 0:
				curLen = 0
			}
		}
	}

	return maxRunLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
