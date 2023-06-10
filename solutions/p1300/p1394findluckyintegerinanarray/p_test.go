package p1394findluckyintegerinanarray

func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, x := range arr {
		m[x]++
	}
	res := -1
	for val, count := range m {
		if count == val && val > res {
			res = val
		}
	}
	return res
}
