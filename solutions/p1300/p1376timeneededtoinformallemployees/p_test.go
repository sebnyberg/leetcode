package p1376timeneededtoinformallemployees

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	// Can just do DFS and find maximum time.
	emp := make([][]int, n)
	for i := range manager {
		if manager[i] != -1 {
			emp[manager[i]] = append(emp[manager[i]], i)
		}
	}
	res := dfs(emp, informTime, headID)
	return res
}

func dfs(emp [][]int, informTime []int, i int) int {
	if len(emp[i]) == 0 {
		return 0
	}
	var res int
	for _, x := range emp[i] {
		res = max(res, dfs(emp, informTime, x))
	}
	return informTime[i] + res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
