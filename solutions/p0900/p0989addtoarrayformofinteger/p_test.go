package p0989addtoarrayformofinteger

func addToArrayForm(num []int, k int) []int {
	rev := func(arr []int) {
		for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	res := make([]int, len(num), len(num)+4)
	copy(res, num)
	rev(res)
	var carry int
	var i int
	for carry > 0 || k > 0 {
		if i >= len(num) {
			res = append(res, 0)
		}
		res[i] += k%10 + carry
		carry = res[i] / 10
		res[i] %= 10
		k /= 10
		i++
	}
	rev(res)
	return res
}
