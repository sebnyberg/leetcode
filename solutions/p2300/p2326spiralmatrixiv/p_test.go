package p2326spiralmatrixiv

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}
	curr := head
	var i, j int
	var k int
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && res[i][j] == -1
	}
	for {
		res[i][j] = curr.Val
		curr = curr.Next
		if curr == nil {
			break
		}
		ii := i + dirs[k][0]
		jj := j + dirs[k][1]
		for !ok(ii, jj) {
			k = (k + 1) % 4
			ii = i + dirs[k][0]
			jj = j + dirs[k][1]
		}
		i = ii
		j = jj
	}
	return res
}
