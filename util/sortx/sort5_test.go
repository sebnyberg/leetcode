package sortx

import (
	"sort"
	"testing"
)

var res []int

func BenchmarkSort(b *testing.B) {
	l := [5]int{5, 9, 2, 3, 5}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		cpy := l[:]
		b.StartTimer()
		MergeSort(cpy)
	}
}

func BenchmarkSortBuiltin(b *testing.B) {
	l := [5]int{5, 9, 2, 3, 5}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		cpy := l[:]
		b.StartTimer()
		sort.Ints(cpy)
	}
}
