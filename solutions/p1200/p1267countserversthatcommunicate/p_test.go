package p1267countserversthatcommunicate

func countServers(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	rows := make([][][]int, m)
	cols := make([][][]int, n)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			rows[i] = append(rows[i], []int{i, j})
			cols[j] = append(cols[j], []int{i, j})
		}
	}

	var res int

	curr := [][]int{}
	next := [][]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			count := 1
			curr = append(curr[:0], []int{i, j})
			grid[i][j] = 0
			for len(curr) > 0 {
				next = next[:0]
				for _, x := range curr {
					for _, srv := range rows[x[0]] {
						if grid[srv[0]][srv[1]] == 0 {
							continue
						}
						grid[srv[0]][srv[1]] = 0
						count++
						next = append(next, srv)
					}
					for _, srv := range cols[x[1]] {
						if grid[srv[0]][srv[1]] == 0 {
							continue
						}
						grid[srv[0]][srv[1]] = 0
						count++
						next = append(next, srv)
					}
				}
				curr, next = next, curr
			}
			if count > 1 {
				res += count
			}
		}
	}
	return res
}
