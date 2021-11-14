package p0739dailytemperatures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_dailyTemperatures(t *testing.T) {
	for _, tc := range []struct {
		temperatures []int
		want         []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.temperatures), func(t *testing.T) {
			require.Equal(t, tc.want, dailyTemperatures(tc.temperatures))
		})
	}
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := make([]int, 1, 16)
	stack[0] = n - 1
	for i := n - 2; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}

func dailyTemperaturesConstantSpace(temperatures []int) []int {
	var tempIdx [101]int
	n := len(temperatures)
	tempIdx[temperatures[n-1]] = n - 1
	ans := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		temp := temperatures[i]
		for aboveTemp := temp + 1; aboveTemp <= 100; aboveTemp++ {
			delta := tempIdx[aboveTemp] - i
			switch {
			case tempIdx[aboveTemp] == 0:
			case ans[i] == 0 || delta < ans[i]:
				ans[i] = delta
			}
		}
		tempIdx[temp] = i
	}
	return ans
}

var _ = dailyTemperaturesConstantSpace
