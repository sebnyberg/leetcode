package p1710maxunitsonatruck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumUnits(t *testing.T) {
	for _, tc := range []struct {
		boxTypes  [][]int
		truckSize int
		want      int
	}{
		{[][]int{{1, 3}, {2, 2}, {3, 1}}, 4, 8},
		{[][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10, 91},
	} {
		t.Run(fmt.Sprintf("%+v", tc.boxTypes), func(t *testing.T) {
			require.Equal(t, tc.want, maximumUnits(tc.boxTypes, tc.truckSize))
		})
	}
}

const maxBoxSize = 1000

func maximumUnits(boxTypes [][]int, truckSize int) int {
	var unitsPerSize [maxBoxSize + 1]int
	for _, boxType := range boxTypes {
		unitsPerSize[boxType[1]] += boxType[0]
	}
	var totalUnits int
	for size := maxBoxSize; size > 0; size-- {
		totalUnits += unitsPerSize[size] * size
		truckSize -= unitsPerSize[size]
		if truckSize <= 0 {
			totalUnits -= -truckSize * size
			break
		}
	}
	return totalUnits
}
