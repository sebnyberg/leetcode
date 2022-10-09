package p2433findtheoriginalarrayofprefixxor

func findArray(pref []int) []int {
	n := len(pref)
	res := make([]int, n)
	var x int
	for i := range res {
		res[i] = x ^ pref[i]
		x = pref[i]
	}
	return res
}
