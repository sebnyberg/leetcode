package p1237findpositiveintegersolutionforagivenequation

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	// Since the function does something with both x and y, it is guaranteed
	// that it grows by at least 1 each time x or y increases.
	// The value space has its maximum value at x=1000 and y=1000.
	// We could fix the value of x and binary search on y, then check if a
	// solution exists.
	hasZ := func(x int) int {
		lo, hi := 1, 1000
		for lo < hi {
			mid := lo + (hi-lo)/2
			val := customFunction(x, mid)
			if val == z {
				return mid
			}
			if val < z {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		return -1
	}
	var res [][]int
	for x := 1; x <= 1000; x++ {
		if y := hasZ(x); y != -1 {
			res = append(res, []int{x, y})
		}
	}
	return res
}
