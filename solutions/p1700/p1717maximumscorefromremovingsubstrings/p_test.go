package p1717maximumscorefromremovingsubstrings

func maximumGain(s string, x int, y int) int {
	// The main question is whether there is a greedy heuristic that always work
	// Let's start with the assumption that the maximum point reduction is optimal
	// Can we come up with a scenario where this is not true?
	//
	// For a reduction to be ambiguous (at least until proven otherwise), there
	// must exist a sequence of "aba". The question is: can such a reduction that
	// does not maximize the one-step value cause a situation where the total
	// value is not maximized?
	//
	// Let's consider the cases
	// xabax
	// A left reduction leads to xax
	// A right reduction leads to xax
	// This means that no matter which action is chosen, the result is the same.
	// If further reductions are possible after, then they were also possible
	// prior.
	//
	// This means that the total number of reductions possible is the same no
	// matter the order, which in turn means that we can greedily reduce according
	// to the maximum options.

	// To make processing easier, let's prioritize left-reductions. If y > x, then
	// we can reverse the string and swap x,y for the same outcome.
	if y > x {
		x, y = y, x
		bs := []byte(s)
		for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
			bs[l], bs[r] = bs[r], bs[l]
		}
		s = string(bs)
	}

	// Now we may greedily reduce any "ab" immediately in the first round. Then
	// any "ba" in the second round
	stack := []byte{}
	var result int
	for _, ch := range s {
		stack = append(stack, byte(ch))
		for len(stack) > 1 && string(stack[len(stack)-2:]) == "ab" {
			stack = stack[:len(stack)-2]
			result += x
		}
	}

	// Reset
	s = string(stack)
	stack = stack[:0]

	// Reduce "ba"
	for _, ch := range s {
		stack = append(stack, byte(ch))
		for len(stack) > 1 && string(stack[len(stack)-2:]) == "ba" {
			stack = stack[:len(stack)-2]
			result += y
		}
	}

	return result
}
