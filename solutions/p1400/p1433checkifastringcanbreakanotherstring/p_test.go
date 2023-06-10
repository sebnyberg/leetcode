package p1433checkifastringcanbreakanotherstring

import (
	"sort"
)

func checkIfCanBreak(s1 string, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	a := true
	b := true
	for i := range s1 {
		a = a && b1[i] >= b2[i]
		b = b && b2[i] >= b1[i]
	}
	return a || b
}
