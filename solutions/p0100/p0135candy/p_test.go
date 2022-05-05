package p0135candy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_candy(t *testing.T) {
	for _, tc := range []struct {
		ratings []int
		want    int
	}{
		{[]int{1, 2, 2}, 4},
		{[]int{1, 0, 2}, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.ratings), func(t *testing.T) {
			require.Equal(t, tc.want, candy(tc.ratings))
		})
	}
}

func candy(ratings []int) int {
	// Two passes: first pass, from left to right
	// ensure that each child has 1 candy, and that
	// each child with a higher rating than the previous has +1 from the previous
	candy := make([]int, len(ratings))
	candy[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		} else {
			candy[i] = 1
		}
	}

	// second pass: right-to-left
	sum := candy[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candy[i] = max(candy[i], candy[i+1]+1)
		}
		sum += candy[i]
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
