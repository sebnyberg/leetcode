package p1053previouspermutationwithoneswap

func prevPermOpt1(arr []int) []int {
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] <= arr[i+1] {
			continue
		}
		idx := i + 1
		val := arr[i+1]
		for j := i + 2; j < len(arr); j++ {
			if arr[i] > arr[j] && arr[j] > val {
				val = arr[j]
				idx = j
			}
		}
		// Find the largest element smaller than the current
		arr[i], arr[idx] = arr[idx], arr[i]
		break
	}
	return arr
}
