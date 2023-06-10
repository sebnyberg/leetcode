package p1385findthedistancevaluebetweentwoarrays

import "sort"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr1)
	sort.Ints(arr2)

	var res int
	for _, x := range arr1 {
		i := sort.SearchInts(arr2, x)
		pre := arr2[0]
		post := arr2[0]
		if i < len(arr2) {
			post = arr2[i]
		}
		if i > 0 {
			pre = arr2[i-1]
		}
		if min(abs(x-pre), abs(post-x)) <= d {
			continue
		}
		res++
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
