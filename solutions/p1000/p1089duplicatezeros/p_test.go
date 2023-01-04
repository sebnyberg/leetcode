package p1089duplicatezeros

func duplicateZeros(arr []int) {
	n := len(arr)
	res := make([]int, n)
	var i int
	for j := 0; j < n; j++ {
		res[j] = arr[i]
		if arr[i] == 0 && j < n-1 {
			j++
			res[j] = 0
		}
		i++
	}
	for i := range res {
		arr[i] = res[i]
	}
}
