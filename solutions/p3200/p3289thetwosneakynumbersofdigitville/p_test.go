package p3289thetwosneakynumbersofdigitville

func getSneakyNumbers(nums []int) []int {
	m := make(map[int]struct{})
	var res []int
	for _, x := range nums {
		if _, exists := m[x]; exists {
			res = append(res, x)
		}
		m[x] = struct{}{}
	}
	return res
}
