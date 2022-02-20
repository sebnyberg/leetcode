package p2183countarraypairsdivisiblebyk

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPairs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{8, 10, 2, 5, 9, 6, 3, 8, 2}, 6, 18},
		{[]int{2, 2, 3, 5, 6, 8, 8, 9, 10}, 6, 18},
		{[]int{1, 2, 3, 4, 5}, 2, 7},
		{[]int{1, 2, 3, 4}, 5, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countPairs(tc.nums, tc.k))
		})
	}
}

func countPairs(nums []int, k int) int64 {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	n := int(math.Sqrt(float64(k)))

	gcds := make([]int, 0, n)
	gcdIdx := make(map[int]int, n)
	for x := 1; x*x <= k; x++ {
		if k%x != 0 {
			continue
		}
		gcdIdx[x] = len(gcds)
		gcds = append(gcds, x)
		y := k / x
		if y == x {
			continue
		}
		gcdIdx[y] = len(gcds)
		gcds = append(gcds, y)
	}
	counts := make([]int, len(gcds))
	for _, num := range nums {
		counts[gcdIdx[gcd(num, k)]]++
	}

	var res int
	for i, a := range gcds {
		n1 := counts[i]
		if (a*a)%k == 0 { // don't forget same-number pairs
			res += n1 * (n1 - 1) / 2
		}
		for j := i + 1; j < len(gcds); j++ {
			n2 := counts[j]
			b := gcds[j]
			if (a*b)%k == 0 {
				res += n1 * n2
			}
		}
	}
	return int64(res)
}
