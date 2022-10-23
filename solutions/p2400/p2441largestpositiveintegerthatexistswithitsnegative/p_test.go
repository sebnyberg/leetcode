package p2441largestpositiveintegerthatexistswithitsnegative

func findMaxK(nums []int) int {
	m := make(map[int]struct{})
	for _, x := range nums {
		m[x] = struct{}{}
	}
	res := -1
	for v := range m {
		if _, exists := m[-v]; exists {
			if v > res {

				res = v
			}
		}
	}
	return res
}
