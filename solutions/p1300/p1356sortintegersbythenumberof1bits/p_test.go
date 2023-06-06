package p1356sortintegersbythenumberof1bits

import (
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a := bits.OnesCount(uint(arr[i]))
		b := bits.OnesCount(uint(arr[j]))
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})
	return arr
}
