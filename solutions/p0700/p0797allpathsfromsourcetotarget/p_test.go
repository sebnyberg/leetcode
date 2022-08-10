package p0797allpathsfromsourcetotarget

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	pre := make([]int, 0, n)
	var res [][]int
	var findPaths func(pre []int, i int)
	findPaths = func(pre []int, i int) {
		pre = append(pre, i)
		if i == n-1 {
			cpy := make([]int, len(pre))
			copy(cpy, pre)
			res = append(res, cpy)
			return
		}
		for _, nei := range graph[i] {
			findPaths(pre, nei)
		}
	}
	findPaths(pre, 0)
	return res
}
