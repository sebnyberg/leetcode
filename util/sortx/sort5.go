package sortx

func MergeSort(arr []int) {
	buf := make([]int, len(arr))
	// merge merges left and right, returning the result
	merge := func(left, right []int) []int {
		buf = buf[:0]
		var l, r int
		nl, nr := len(left), len(right)
		for l < nl || r < nr {
			if l == nl {
				buf = append(buf, right[r:]...)
				break
			} else if r == nr {
				buf = append(buf, left[l:]...)
				break
			} else if left[l] <= right[r] {
				buf = append(buf, left[l])
				l++
			} else {
				buf = append(buf, right[r])
				r++
			}
		}
		return buf
	}

	var msort func([]int)
	msort = func(subarr []int) {
		if len(subarr) <= 1 {
			return
		}
		mid := len(subarr) / 2
		msort(subarr[:mid])
		msort(subarr[mid:])
		tmp := merge(subarr[:mid], subarr[mid:])
		copy(subarr, tmp)
	}

	msort(arr)
}
