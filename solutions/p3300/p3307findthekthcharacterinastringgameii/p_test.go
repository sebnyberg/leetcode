package p3307findthekthcharacterinastringgameii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthCharacter(t *testing.T) {
	for _, tc := range []struct {
		k          int64
		operations []int
		want       byte
	}{
		{3, []int{1, 0}, 'a'},
		{1, []int{0, 0, 0}, 'a'},
		{10, []int{0, 1, 0, 1}, 'b'},
		{5, []int{0, 0, 0}, 'a'},
		{2, []int{1}, 'b'},
	} {
		t.Run(fmt.Sprintf("%+v", tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, kthCharacter(tc.k, tc.operations))
		})
	}
}

func kthCharacter(k int64, operations []int) byte {
	bitsNeeded := 1
	for kk := k - 1; kk > 0; kk /= 2 {
		bitsNeeded++
	}
	operations = operations[:bitsNeeded-1]
	val := f(int(k-1), 1<<(len(operations)), operations)
	return byte('a' + val%26)
}

func f(k int, n int, operations []int) int {
	if n == 1 {
		return 0
	}
	toRemove := 1 << (len(operations) - 1)
	n -= toRemove
	currentOp := operations[len(operations)-1]
	operations = operations[:len(operations)-1]
	if k >= n {
		kk := k - n
		return currentOp + f(kk, n, operations)
	}
	return f(k, n, operations)
}
