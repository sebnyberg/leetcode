package p1304findnuniqueintegerssumuptozero

func sumZero(n int) []int {
	var res []int
	for i := 0; i < n/2; i++ {
		res = append(res, i+1)
		res = append(res, -(i + 1))
	}
	if n%2 == 1 {
		res = append(res, 0)
	}
	return res
}
