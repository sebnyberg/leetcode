package p1826faultysensor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_badSensor(t *testing.T) {
	for _, tc := range []struct {
		sensor1 []int
		sensor2 []int
		want    int
	}{
		{[]int{2, 3, 4, 5}, []int{2, 1, 3, 4}, 1},
		{[]int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 5}, -1},
		{
			[]int{2, 3, 2, 2, 3, 2},
			[]int{2, 3, 2, 3, 2, 7},
			2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.sensor1), func(t *testing.T) {
			require.Equal(t, tc.want, badSensor(tc.sensor1, tc.sensor2))
		})
	}
}

func badSensor(sensor1 []int, sensor2 []int) int {
	n := len(sensor1)
	for i := 0; i < n-1; i++ {
		if sensor1[i] == sensor2[i] {
			continue
		}
		for ; i < n-1; i++ {
			if sensor1[i] != sensor2[i+1] {
				return 2
			}
			if sensor2[i] != sensor1[i+1] {
				return 1
			}
		}
	}
	return -1
}
