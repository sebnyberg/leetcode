package p2931maximumspendingafterbuyingitems

import "sort"

func maxSpending(values [][]int) int64 {
	// Hm, isn't this just a matter of buying products from cheap to expensive?
	// We could just sort all values...
	var concat []int
	for _, v := range values {
		concat = append(concat, v...)
	}
	sort.Ints(concat)
	var res int
	for i := range concat {
		res += (i + 1) * concat[i]
	}

	return int64(res)
}
