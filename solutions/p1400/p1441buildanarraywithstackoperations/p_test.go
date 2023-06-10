package p1441buildanarraywithstackoperations

func buildArray(target []int, n int) []string {
	x := 1
	var res []string
	for i := range target {
		for x < target[i] {
			res = append(res, "Push", "Pop")
			x++
		}
		res = append(res, "Push")
		x++
	}
	return res
}
