package p0324wigglesort2

// func Test_wiggleSort(t *testing.T) {
// 	for _, tc := range []struct {
// 		nums []int
// 		want []int
// 	}{
// 		{[]int{1, 5, 1, 1, 6, 4}, []int{1, 6, 1, 5, 1, 4}},
// 		{[]int{1, 3, 2, 2, 3, 1}, []int{2, 3, 1, 3, 1, 2}},
// 	} {
// 		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
// 			wiggleSort(tc.nums)
// 			require.Equal(t, tc.want, tc.nums)
// 		})
// 	}
// }

// // TODO
// func wiggleSort(nums []int) {
// 	n := len(nums)
// 	var numCounts [5001]int
// 	for _, n := range nums {
// 		numCounts[n]++
// 	}
// 	var counts int
// 	var lower, lowerCount int
// 	for num := 0; num <= 5000; num++ {
// 		diff := counts + numCounts[num] - n/2
// 		if diff >= 0 {
// 			// we have found the median
// 			lower = num
// 			lowerCount = diff
// 			numCounts[num] -= diff
// 			break
// 		}
// 		counts += numCounts[num]
// 	}
// 	upper := 5000
// }
