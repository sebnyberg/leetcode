package p2357makearrayzerobysubtractingequalamounts

func minimumOperations(nums []int) int {
	m := make(map[int]struct{})
	for _, x := range nums {
		m[x] = struct{}{}
	}
	res := len(m)
	if _, exists := m[0]; exists {
		res--
	}
	return res
}
