package p1278palindromepartitioningiii

import "math"

type state struct {
	s string
	k int
}

func palindromePartition(s string, k int) int {
	// Perhaps the problem can be split into multiple sub-problems.
	// For example, it is possible to calculate the minimum amount of character
	// swaps needed to construct a palindrome between two indices in the string.
	// Then we would need to combine segments of the string so as to minimize
	// the total number of swaps, which sounds like a knapsack-type problem.
	//
	// Another way to look at it is recursively: we could find the minimum cost
	// of f(s, 2) by findinng the minimum combination of f(s[:i], 1) + f(s[i:],
	// k-1) for all i < n-k.
	//
	// Please note that because strings are immutable, storing a string in a map
	// is the same as storing a string header, which is just a length and a data
	// pointer (usually 128bits). Having a string as part of the key therefore
	// has the same footprint as having a start/end offset interval.
	mem := make(map[state]int)

	return f(mem, s, k)
}

func f(mem map[state]int, s string, k int) int {
	key := state{s, k}
	if v, exists := mem[key]; exists {
		return v
	}
	if len(s) == 1 {
		return 0
	}
	if k == 1 {
		// calculate minimum cost of forming palindrome from this string
		var cost int
		for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				cost++
			}
		}
		return cost
	}

	// Try combinations of segments to find best solution
	res := math.MaxInt32
	for i := 1; i <= len(s)-(k-1); i++ {
		res = min(res, f(mem, s[:i], 1)+f(mem, s[i:], k-1))
	}

	mem[key] = res
	return mem[key]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
