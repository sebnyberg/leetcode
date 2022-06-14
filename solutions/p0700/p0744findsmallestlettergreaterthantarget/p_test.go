package p1

import "sort"

func nextGreatestLetter(letters []byte, target byte) byte {
	i := sort.Search(len(letters), func(i int) bool {
		return letters[i] > target
	})
	if i == len(letters) {
		return letters[0]
	}
	return letters[i]
}
