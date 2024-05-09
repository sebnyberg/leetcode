package p3075maximizehappinessofselectedchildren

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	// Since the change in happiness is independent of the position of the
	// children, the order of picked children is irrelevant. Therefore, we can
	// just pick children in order of size and decrement by i each time until k ==
	// 0 or the current happiness is 0.
	sort.Sort(sort.Reverse(sort.IntSlice(happiness)))
	var res int
	for i := 0; i < len(happiness) && k > 0 && happiness[i]-i > 0; i++ {
		res += happiness[i] - i
		k--
	}
	return int64(res)
}
