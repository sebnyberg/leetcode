package p1860incrementalmemoryleak

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_memLeak(t *testing.T) {
	for _, tc := range []struct {
		memory1 int
		memory2 int
		want    []int
	}{
		{2, 2, []int{3, 1, 0}},
		{8, 11, []int{6, 0, 4}},
		{int(10e9), 1, []int{6, 0, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.memory1), func(t *testing.T) {
			require.Equal(t, tc.want, memLeak(tc.memory1, tc.memory2))
		})
	}
}

func memLeak(memory1 int, memory2 int) []int {
	// orig := [2]int{memory1, memory2}
	mem := [2]int{memory1, memory2}
	i := 1
	for {
		maxIdx := 0
		if mem[1] > mem[0] {
			maxIdx = 1
		}
		if mem[maxIdx]-i < 0 {
			break
		}
		mem[maxIdx] -= i
		i++
	}
	return []int{i, mem[0], mem[1]}
}
