package p2818applyoperationstomaximizescore

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumScore(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{8, 3, 9, 3, 8}, 2, 81},
		{[]int{19, 12, 14, 6, 10, 18}, 3, 4788},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumScore(tc.nums, tc.k))
		})
	}
}

const mod = 1e9 + 7

func maximumScore(nums []int, k int) int {
	n := len(nums)

	// Precompute prime scores for all numbers up to the maximum in nums
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	primeScores := make([]int, maxNum+1)
	for i := 2; i <= maxNum; i++ {
		if primeScores[i] == 0 { // i is prime
			for j := i; j <= maxNum; j += i {
				primeScores[j]++
			}
		}
	}

	// Compute next greater elements and previous greater elements for each element
	nextGreater := make([]int, n)
	prevGreater := make([]int, n)

	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && primeScores[nums[stack[len(stack)-1]]] < primeScores[nums[i]] {
			nextGreater[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		nextGreater[stack[len(stack)-1]] = n
		stack = stack[:len(stack)-1]
	}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && primeScores[nums[stack[len(stack)-1]]] <= primeScores[nums[i]] {
			prevGreater[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		prevGreater[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	// Create elements with their contribution counts
	type Element struct {
		val   int
		count int
	}

	elements := make([]Element, n)
	for i := 0; i < n; i++ {
		left := prevGreater[i]
		right := nextGreater[i]
		count := (i - left) * (right - i)
		elements[i] = Element{nums[i], count}
	}

	// Sort elements in descending order
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].val > elements[j].val
	})

	// Calculate the maximum score
	score := 1
	remaining := k

	for _, elem := range elements {
		if remaining <= 0 {
			break
		}
		take := elem.count
		if take > remaining {
			take = remaining
		}
		score = (score * pow(elem.val, take)) % mod
		remaining -= take
	}

	return score
}

func pow(x, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = (res * x) % mod
		}
		x = (x * x) % mod
		n /= 2
	}
	return res
}
